package dm

import (
	"vesta_mlp/util/database"
	"vesta_mlp/util/queue"
)

//db和文件服务器配置信息
type Struct_Config struct {
	DBInfo      *database.GormDBInfo
	FileUpLoad  *FileUpLoadURL
	TokenParams *TokenRequest
	Rmq         *queue.RMQ
	Pro         bool
}

type FileUpLoadURL struct {
	UrlToken string
	UrlUp    string
	UrlLoad  string
}

type TokenRequest struct {
	BusiType     string `json:"busiType"`     //"busiType":"order",
	BusiId       string `json:"busiId"`       //"busiId":"111111",
	TemplateCode string `json:"templateCode"` //"templateCode":"yixin_xiaofei_tibao_polaris",
	CreatorId    string `json:"creatorId"`    //"creatorId":"zhangsan",
	CreatorName  string `json:"creatorName"`  //"creatorName":"张三"
}
