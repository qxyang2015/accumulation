package parse_syntax

//
//import (
//	"fmt"
//	"github.com/qxyang2015/accumulation/tools/common"
//	dm "github.com/qxyang2015/accumulation/web/datamanager"
//	"strconv"
//	"strings"
//)
//
////表达式建立树形结构
//func ConstructSyntaxTree(express []ExpressNode) ([]SyntaxTreeNode, error) {
//	syntaxTreeLen := len(express)
//	syntaxTree := make([]SyntaxTreeNode, syntaxTreeLen)
//	if len(express) == 0 {
//		return syntaxTree, nil
//	}
//	stack := common.NewStack()
//	for idx, expNode := range express {
//		//S := expNode.Value
//		syntaxTreeNode := SyntaxTreeNode{}
//		if expNode.Type == base_enum.ExpressOperator {
//			syntaxTreeNode.Index = len(express) - idx - 1
//			syntaxTreeNode.PNode = -1
//			if _, ok := dm.Config.OperatorMap[expNode.Value]; !ok {
//				log4sys.Error("OperatorMap operator[%s] is not found", expNode.Value)
//				return nil, fmt.Errorf("OperatorMap operator[%s] is not found", expNode.Value)
//			}
//			if dm.Config.OperatorMap[expNode.Value].ValueNum == 1 {
//				stackPrevVal := stack.Pop()
//				if stackPrevVal == nil {
//					log4sys.Error("stack is empty operator clculate is null")
//					return nil, fmt.Errorf("stack is empty operator clculate is null")
//				}
//				prevVal := stackPrevVal.(model.SyntaxTreeNode)
//				syntaxTreeNode.CNodes = append(syntaxTreeNode.CNodes, prevVal.Index)
//				prevVal.PNode = syntaxTreeNode.Index
//				//ResultTree = append(ResultTree, prevVal)
//				syntaxTree[prevVal.Index] = prevVal
//			} else if _, ok := dm.Config.OperatorMap[expNode.Value]; ok {
//				stackBackVal := stack.Pop()
//				stackPrevVal := stack.Pop()
//				if stackBackVal == nil || stackPrevVal == nil {
//					log4sys.Error("stack is empty operate clculate is null")
//					return nil, fmt.Errorf("stack is empty operate clculate is null")
//				}
//				backVal := stackBackVal.(model.SyntaxTreeNode)
//				prevVal := stackPrevVal.(model.SyntaxTreeNode)
//				syntaxTreeNode.CNodes = append(syntaxTreeNode.CNodes, prevVal.Index)
//				syntaxTreeNode.CNodes = append(syntaxTreeNode.CNodes, backVal.Index)
//				prevVal.PNode = syntaxTreeNode.Index
//				backVal.PNode = syntaxTreeNode.Index
//				//ResultTree = append(ResultTree, prevVal)
//				//ResultTree = append(ResultTree, backVal)
//				syntaxTree[prevVal.Index] = prevVal
//				syntaxTree[backVal.Index] = backVal
//			} else {
//				log4sys.Error("operate is not exit[%s]", expNode.Value)
//				return nil, fmt.Errorf("operate is not exit[%s]", expNode.Value)
//			}
//			if expNode.Type == base_enum.ExpressOperator {
//				syntaxTreeNode.Type = expNode.Type
//				syntaxTreeNode.ValType = expNode.ValType
//			} else {
//				log4sys.Debug("ConstructTree not found FieldType[Json2Tree]:%s", express[idx].Type)
//			}
//			syntaxTreeNode.Value = expNode.Value
//			stack.Push(syntaxTreeNode)
//		} else if expNode.Type == base_enum.ExpressConstant || expNode.Type == base_enum.ExpressVariable || expNode.Type == base_enum.ExpressFeature {
//			syntaxTreeNode.Index = len(express) - idx - 1
//			//值类型:bool/number/string/array_n/array_s,操作符都为string类型
//			syntaxTreeNode.Type = expNode.Type
//			syntaxTreeNode.ValType = expNode.ValType
//			if syntaxTreeNode.Type == base_enum.ExpressConstant {
//				//转换函数
//				v, err := TypeTransform(expNode.Value, expNode.ValType)
//				if err != nil {
//					log4sys.Error("TypeTransform error[%v]", err)
//					return nil, fmt.Errorf("TypeTransform error[%v]", err)
//				}
//				syntaxTreeNode.Value = v
//			} else {
//				syntaxTreeNode.Value = expNode.Value
//			}
//			stack.Push(syntaxTreeNode)
//		} else if expNode.Type == base_enum.ExpressFunc {
//			//fmt.Println(exp.Value)
//			syntaxTreeNode.Index = len(express) - idx - 1
//			//函数类型：Type转换为操作符类型
//			syntaxTreeNode.Type = expNode.Type
//			syntaxTreeNode.ValType = expNode.ValType
//			syntaxTreeNode.Value = expNode.Value
//			cnodes := make([]int, len(expNode.Param))
//			for parmNum, paramNode := range expNode.Param {
//				paramSyntaxTreeNode := model.SyntaxTreeNode{
//					Index:   syntaxTreeLen + parmNum,
//					PNode:   syntaxTreeNode.Index,
//					Type:    paramNode.ParamType,
//					ValType: paramNode.ValType,
//					Value:   paramNode.Value,
//				}
//				if paramSyntaxTreeNode.Type == base_enum.ExpressConstant {
//					v, err := TypeTransform(paramNode.Value, paramNode.ValType)
//					if err != nil {
//						log4sys.Error("TypeTransform is error[%v]", err)
//						return nil, fmt.Errorf("TypeTransform is error[%v]", err)
//					}
//					paramSyntaxTreeNode.Value = v
//				}
//				cnodes[parmNum] = paramSyntaxTreeNode.Index
//				syntaxTree = append(syntaxTree, paramSyntaxTreeNode)
//			}
//			syntaxTreeNode.CNodes = cnodes
//			syntaxTreeLen += len(expNode.Param)
//			stack.Push(syntaxTreeNode)
//		} else {
//			log4sys.Error("fieldType[%s] not operator,constant,variable", expNode.Type)
//			return nil, fmt.Errorf("fieldType[%s] not operator,constant,variable", expNode.Type)
//		}
//	}
//	//ResultTree = append(ResultTree, stack.Pop().(SyntaxTreeNode))
//	sytaxTreeNode := stack.Pop().(model.SyntaxTreeNode)
//	syntaxTree[sytaxTreeNode.Index] = sytaxTreeNode
//	if stack.Len() != 0 {
//		log4sys.Error("ConstructTree stack len[%v] is not nil", stack.Len())
//		return nil, fmt.Errorf("ConstructTree stack len[%v] is not nil", stack.Len())
//	}
//	return syntaxTree, nil
//}
//
////填充表达式内值类型
////非函数类型：变量、特征、操作符：文本类型；常量：数值类型
////函数类型：依照函数配置
//func ExpressFillValType(express []model.ExpressNode, funcMap map[string]model.VstEditFuncDtl) ([]model.ExpressNode, error) {
//	log4sys := logging.GetLogger()
//	log4sys.Trace("ExpressFillValType start")
//	for idx, expNode := range express {
//		//非函数类型
//		if expNode.Type != base_enum.ExpressFunc {
//			if expNode.Type == base_enum.ExpressOperator {
//				express[idx].ValType = base_enum.VTypeString
//			} else if expNode.Type == base_enum.ExpressVariable {
//				if v, ok := dm.Config.VariableMap[expNode.Value]; ok {
//					express[idx].ValType = v.DataType
//				} else {
//					log4sys.Error("variable id[%v] variableMap is nil", expNode.Value)
//					return nil, fmt.Errorf("variable id[%v] variableMap is nil", expNode.Value)
//				}
//			} else if expNode.Type == base_enum.ExpressFeature {
//				if f, ok := dm.Config.FeatureMap[expNode.Value]; ok {
//					express[idx].ValType = f.DataType
//				} else {
//					log4sys.Error("feature id[%v] featureMap is nil", expNode.Value)
//					return nil, fmt.Errorf("feature id[%v] featureMap is nil", expNode.Value)
//				}
//			} else if expNode.Type == base_enum.ExpressConstant {
//				express[idx].ValType = base_enum.VTypeNumber
//			}
//		} else {
//			if funcVal, ok := funcMap[expNode.Value]; ok {
//				express[idx].ValType = funcVal.OutputParam
//				for pIdx, _ := range expNode.Param {
//					//格式校验，保存时已经校验
//					expNode.Param[pIdx].ValType = funcVal.InputParamList[pIdx].ParamType
//				}
//			}
//		}
//	}
//	return express, nil
//}
//
////中缀表达式转后缀表达式
////遇到数字直接输出
////遇到运算符则判断：
////栈顶运算符优先级更低则入栈，更高或相等则直接输出直到遇到优先级低的运算符
////栈为空、栈顶是 ( 直接入栈
////运算符是 ) 则将栈顶运算符全部弹出，直到遇见)
////中缀表达式遍历完毕，运算符栈不为空则全部弹出，依次追加到输出
//func ExpressInToPost(InOrderExp []model.ExpressNode, operatorMap map[string]model.VstOperatorTable) ([]model.ExpressNode, error) {
//	log4sys := logging.GetLogger()
//	log4sys.Trace("ExpressExplain start")
//	var postOrderExp []model.ExpressNode
//	stack := common.NewStack()
//	for _, exp := range InOrderExp {
//		s := exp.Value
//		if exp.Type == base_enum.ExpressOperator {
//			if s == "(" {
//				stack.Push(exp)
//			} else if s == ")" {
//				for {
//					sVal := stack.Pop().(model.ExpressNode)
//					if sVal.Value == "(" {
//						break
//					}
//					postOrderExp = append(postOrderExp, sVal)
//				}
//			} else {
//				//操作符：遇到高优先级的运算符，不断弹出，直到遇见更低优先级运算符
//				for !stack.Empty() {
//					topVal := stack.Peak().(model.ExpressNode).Value
//					if topVal == "(" {
//						break
//					}
//					if _, ok1 := operatorMap[exp.Value]; !ok1 {
//						log4sys.Error("opInfo exp.Value[%s] is not found", exp.Value)
//						return nil, fmt.Errorf("opInfo exp.Value[%s] is not found", exp.Value)
//					}
//					if _, ok2 := dm.Config.OperatorMap[topVal]; !ok2 {
//						log4sys.Error("opInfo topVal[%s] is not found", topVal)
//						return nil, fmt.Errorf("opInfo topVal[%s] is not found", topVal)
//					}
//					if operatorMap[exp.Value].Salience > operatorMap[topVal].Salience {
//						break
//					}
//					postOrderExp = append(postOrderExp, stack.Pop().(model.ExpressNode))
//				}
//				stack.Push(exp)
//			}
//		} else if exp.Type == base_enum.ExpressVariable || exp.Type == base_enum.ExpressConstant || exp.Type == base_enum.ExpressFunc || exp.Type == base_enum.ExpressFeature {
//			postOrderExp = append(postOrderExp, exp)
//		} else {
//			log4sys.Error("fieldType[%s] not operator,constant,variable,func", exp.Type)
//			return nil, fmt.Errorf("fieldType[%s] not operator,constant,variable,func", exp.Type)
//		}
//	}
//	//将栈内元素全部弹出
//	for {
//		if stack.Empty() {
//			break
//		}
//		postOrderExp = append(postOrderExp, stack.Pop().(model.ExpressNode))
//	}
//	return postOrderExp, nil
//}
//
////表达式结构转语法树结构
//func ExpDataToSyntaxTreeData(expressData []model.ExpressNode) []model.SyntaxTreeNode {
//	log4sys := logging.GetLogger()
//	log4sys.Trace("ExpDataToSyntaxTreeData start")
//	syntaxTree := make([]model.SyntaxTreeNode, len(expressData))
//	for idx, exp := range expressData {
//		syntaxTree[idx] = model.SyntaxTreeNode{
//			Index: idx,
//			PNode: idx - 1,
//			Type:  exp.Type,
//			Value: exp.Value,
//		}
//		//syntaxTree[idx].CNodes=append(syntaxTree[idx].CNodes,idx)
//	}
//	log4sys.Trace("ExpDataToSyntaxTreeData end")
//	return syntaxTree
//}
//
////值类型:number/string/array_n/array_s,操作符都为string类型
//func TypeTransform(str string, valType string) (interface{}, error) {
//	log4sys := logging.GetLogger()
//	switch valType {
//	case base_enum.VTypeNumber:
//		val, err := strconv.ParseFloat(str, 64)
//		if err != nil {
//			log4sys.Error("number type error[%s]", str)
//			return nil, fmt.Errorf("number type error[%s]", str)
//		}
//		return val, nil
//	case base_enum.VTypeString:
//		return str, nil
//	case base_enum.VTypeArrNum:
//		var resStr []float64
//		strSlice := strings.Split(str, ",")
//		for _, strVal := range strSlice {
//			val, err := strconv.ParseFloat(strVal, 64)
//			if err != nil {
//				log4sys.Error("array_n type error[%s]", str)
//				return nil, fmt.Errorf("array_n type error[%s]", str)
//			}
//			resStr = append(resStr, val)
//		}
//		return resStr, nil
//	case base_enum.VTypeArrStr:
//		strSlice := strings.Split(str, ",")
//		return strSlice, nil
//	default:
//		log4sys.Error("value type not exit:%s", valType)
//		return nil, fmt.Errorf("value type not exit", valType)
//	}
//}
