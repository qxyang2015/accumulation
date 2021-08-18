package dm

import (
	"github.com/spf13/viper"
	"tools/error_tools"
	"tools/queue"
)

var Config Struct_Config

func InitConf(envName string) {
	viper.SetConfigName("config")            //把json文件换成yaml文件，只需要配置文件名 (不带后缀)即可
	viper.AddConfigPath("./conf/" + envName) //添加配置文件所在的路径
	err := viper.ReadInConfig()
	error_tools.Assert(err)
	if envName == "pro" {
		Config.Pro = true
	}

	//数据库初始化
	dBInfo := db.CreateGormDBByYaml(viper.GetViper(), "mysql.vesta_db")
	Config.DBInfo, err = db.InitGormMysql(dBInfo)
	error_tools.Assert(err)

	Config.FileUpLoad = &FileUpLoadURL{
		UrlToken: viper.GetString("fileupload.urltoken"),
		UrlUp:    viper.GetString("fileupload.urlup"),
		UrlLoad:  viper.GetString("fileupload.urlload"),
	}

	Config.TokenParams = &TokenRequest{
		BusiType:     viper.GetString("fileupload.busitype"),
		BusiId:       viper.GetString("fileupload.busiid"),
		TemplateCode: viper.GetString("fileupload.templatecode"),
		CreatorId:    viper.GetString("fileupload.creatorId"),
		CreatorName:  viper.GetString("fileupload.creatorName"),
	}

	address := viper.GetString("rabbitmq.address")
	exchange := viper.GetString("rabbitmq.exchange")
	routingKey := viper.GetString("rabbitmq.routingkey")

	Config.Rmq, err = queue.InitRMQ(address, exchange, routingKey)
	error_tools.Assert(err)
}
