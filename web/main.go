package main

import (
	"runtime/debug"
	"time"
	"vesta_mlp/util/common"
	"vesta_mlp/util/http_tools"
	"vesta_mlp/util/logging"
	"vesta_mlp/websrv/demo/datamanager"
	"vesta_mlp/websrv/demo/userinterface"
)

const (
	cst_YAMLFILE = "./conf/digg.yaml"
)

func main() {
	startTime := time.Now()

	go func() {
		for {
			debug.FreeOSMemory()
			time.Sleep(1 * time.Minute)
		}
	}()

	argsMap := common.GetCommandLineParam([]string{"config"})

	//执行http接入层Init初始化函数
	cfg := http_tools.WebInit(cst_YAMLFILE, *argsMap["config"])
	log4sys := logging.GetLogger()

	serverInit(*argsMap["config"])

	log4sys.Warn("注册web服务 location=>func")
	ui.Init(cfg.Router)

	log4sys.Warn("服务Init模块加载成功，耗时[%v],Http服务启动", time.Now().Sub(startTime).String())
	http_tools.ServeMain(cfg)
}

//根据不同服务需要，执行不同的初始化操作,注意不同初始化操作执行的先后依赖顺序
func serverInit(envName string) {
	log4sys := logging.GetLogger()
	log4sys.Warn("data manager层 模块初始化开始执行")
	dm.InitConf(envName)

	log4sys.Warn("problem domain层模块初始化开始执行")
	//pd.Init()

}
