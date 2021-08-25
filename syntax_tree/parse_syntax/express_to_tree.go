package parse_syntax

import (
	"encoding/json"
	"fmt"
	dm "github.com/qxyang2015/accumulation/web/datamanager"
)

//表达式节点定义
type ExpressNode struct {
	Type    string      `json:"type"`
	Class   bool        `json:"class"`
	Value   string      `json:"value"`
	ValType string      `json:"valType,omitempty"`
	Label   string      `json:"label"`
	Param   []ParamNode `json:"param,omitempty"`
}

//表达式结构体
type ParamNode struct {
	ParamType        string `json:"paramType"`                  //函数入参类型，number、string、array_n、array_str
	ParamTypeDisplay string `json:"paramTypeDisplay,omitempty"` //函数入参类型展示值
	Value            string `json:"value"`
	ValType          string `json:"valType,omitempty"`
	Name             string `json:"name,omitempty"`
}

//语法树节点定义
type SyntaxTreeNode struct {
	Index   int         `json:"index"`  //节点序号
	PNode   int         `json:"pnode"`  //父节点序号,根节点时为空
	CNodes  []int       `json:"cnodes"` //孩子节点序号列表,当为叶子节点时,该字段为空
	Type    string      `json:"type"`   //字段类型:variable/constant/operator/feature
	ValType string      `json:"vtype"`  //值类型:bool/number/string/array_n/array_s,操作符都为string类型
	Value   interface{} `json:"value"`  //节点值
}

//表达式转语法树
func ExpressToSyntaxTree(express []ExpressNode) ([]SyntaxTreeNode, error) {
	postExpress, err := ExpressInToPost(express, dm.Config.OperatorMap)
	if err != nil {
		log4sys.Error("ExpressInToPost is error[%v]", err)
		return nil, fmt.Errorf("ExpressInToPost is error[%v]", err)
	}
	log4sys.Trace("postExpress:[%v]", postExpress)
	postExpress, err = ExpressFillValType(postExpress, dm.Config.FuncMap)
	if err != nil {
		log4sys.Error("ExpressFillValType is error[%v]", err)
		return nil, fmt.Errorf("ExpressFillValType is error[%v]", err)
	}
	log4sys.Trace("FillValType postExpress:[%v]", postExpress)
	syntaxTree, err := ConstructSyntaxTree(postExpress)
	if err != nil {
		log4sys.Error("ConstructTree is error[%v]", err)
		return nil, fmt.Errorf("ConstructTree is error[%v]", err)
	}
	log4sys.Trace("ConstructTree syntaxTree:[%v]", syntaxTree)
	return syntaxTree, err
}

//表达式转语法树,不构建属性结构
func ExpToSyntaxNoneConstructTree(express []model.ExpressNode) ([]model.SyntaxTreeNode, error) {
	log4sys := logging.GetLogger()
	log4sys.Trace("ExpressToSyntaxTree start")
	expressFill, err := ExpressFillValType(express, dm.Config.FuncMap)
	if err != nil {
		log4sys.Error("ExpressFillValType is error[%v]", err)
		return nil, fmt.Errorf("ExpressFillValType is error[%v]", err)
	}
	syntree := ExpDataToSyntaxTreeData(expressFill)
	log4sys.Trace("ConstructTree syntaxTree:[%v]", syntree)
	return syntree, nil
}

//变量递归获取为列表
//特征递归获取为递归形式
//处理深层引用、循环引用
func GetFeatureQuoteVar(express []model.ExpressNode, varibaleList *[]model.VariableInfo, cnt, maxCnt int) error {
	log4sys := logging.GetLogger()
	cnt++
	db := dm.Config.EditDBInfo
	//限制最大迭代次数
	if cnt > maxCnt {
		log4sys.Error("GetFeatureQuoteVar cnt[%v] > maxCnt[%v]", cnt, maxCnt)
		return fmt.Errorf("GetFeatureQuoteVar cnt[%v] > maxCnt[%v]", cnt, maxCnt)
	}
	for _, expNode := range express {
		if expNode.Type == base_enum.ExpressVariable {
			varCnt := 0
			err := GetVarQuoteVar(expNode.Value, varibaleList, varCnt, utils.Max_Cnt)
			if err != nil {
				log4sys.Error("获取变量引用变量id[%v]出现错误[%v]", expNode.Value, err)
				return fmt.Errorf("获取变量引用变量id[%v]出现错误[%v]", expNode.Value, err)
			}
		} else if expNode.Type == base_enum.ExpressFeature {
			//变量引用变量
			var featureInfo model.VstEditFeatureDtl
			dbRes := db.Table("vst_edit_feature_dtl").Where("id = ?", expNode.Value).First(&featureInfo)
			if dbRes.Error != nil || dbRes.RowsAffected == 0 {
				log4sys.Error("vst_edit_feature_dtl id[%v] query is error[%v] or rowsAffeceted[%v]", expNode.Value, dbRes, dbRes.RowsAffected)
				return fmt.Errorf("vst_edit_feature_dtl id[%v] query is error[%v] or rowsAffeceted[%v]", expNode.Value, dbRes, dbRes.RowsAffected)
			}
			err := json.Unmarshal([]byte(featureInfo.Express), &featureInfo.ExpressList)
			if err != nil {
				log4sys.Error("featureId[%v] json Unmarshal is error[%v]", featureInfo.Id, err)
				return fmt.Errorf("featureId[%v] json Unmarshal is error[%v]", featureInfo.Id, err)
			}
			err = GetFeatureQuoteVar(featureInfo.ExpressList, varibaleList, cnt, maxCnt)
			if err != nil {
				log4sys.Error("GetFeatureQuoteVar featureId[%v] cnt[%v] maxCnt[%v] is error[%v]", expNode.Value, cnt, maxCnt, err)
				return fmt.Errorf("GetFeatureQuoteVar featureId[%v] cnt[%v] maxCnt[%v] is error[%v]", expNode.Value, cnt, maxCnt, err)
			}
		} else if expNode.Type == base_enum.ExpressFunc {
			for _, paramNode := range expNode.Param {
				if paramNode.ParamType == base_enum.ExpressVariable {
					varCnt := 0
					err := GetVarQuoteVar(paramNode.Value, varibaleList, varCnt, maxCnt)
					if err != nil {
						log4sys.Error("func featureId[%v] GetVarQuoteVar is error[%v]", paramNode.Value, err)
						return fmt.Errorf("func featureId[%v] GetVarQuoteVar is error[%v]", paramNode.Value, err)
					}
				} else if paramNode.ParamType == base_enum.ExpressFeature {
					var featureInfo model.VstEditFeatureDtl
					dbRes := db.Table("vst_edit_feature_dtl").Where("id = ?", paramNode.Value).First(&featureInfo)
					if dbRes.Error != nil || dbRes.RowsAffected == 0 {
						log4sys.Error("func vst_edit_feature_dtl id[%v] query is error[%v] or rowsAffeceted[%v]", paramNode.Value, dbRes, dbRes.RowsAffected)
						return fmt.Errorf("func vst_edit_feature_dtl id[%v] query is error[%v] or rowsAffeceted[%v]", paramNode.Value, dbRes, dbRes.RowsAffected)
					}
					err := json.Unmarshal([]byte(featureInfo.Express), &featureInfo.ExpressList)
					if err != nil {
						log4sys.Error("func featureId[%v] json Unmarshal is error[%v]", featureInfo.Id, err)
						return fmt.Errorf("func featureId[%v] json Unmarshal is error[%v]", featureInfo.Id, err)
					}
					err = GetFeatureQuoteVar(featureInfo.ExpressList, varibaleList, cnt, maxCnt)
					if err != nil {
						log4sys.Error("func GetFeatureQuoteVar featureId[%v] cnt[%v] maxCnt[%v] is error[%v]", paramNode.Value, cnt, maxCnt, err)
						return fmt.Errorf("func GetFeatureQuoteVar featureId[%v] cnt[%v] maxCnt[%v] is error[%v]", paramNode.Value, cnt, maxCnt, err)
					}
				}
			}
		}
	}
	return nil
}
