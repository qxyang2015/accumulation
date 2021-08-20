package dm

import (
	"github.com/qxyang2015/accumulation/tools/error_tools"
	"github.com/spf13/viper"
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
	//TODO 配置读取逻辑
	error_tools.Assert(err)
}
