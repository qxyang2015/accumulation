package main

import (
	"fmt"
	"github.com/qxyang2015/accumulation/tools/common"
	"github.com/qxyang2015/accumulation/tools/http_tools"
	dm "github.com/qxyang2015/accumulation/web/datamanager"
	ui "github.com/qxyang2015/accumulation/web/userinterface"
	"runtime/debug"
	"time"
)

const (
	cst_YAMLFILE = "./conf/digg.yaml"
)

func main() {
	fmt.Println("Run Web")
	go func() {
		for {
			debug.FreeOSMemory()
			time.Sleep(1 * time.Minute)
		}
	}()

	argsMap := common.GetCommandLineParam([]string{"config"})

	//执行http接入层Init初始化函数
	cfg := http_tools.WebInit(cst_YAMLFILE, *argsMap["config"])

	serverInit(*argsMap["config"])

	ui.Init(cfg.Router)

	http_tools.ServeMain(cfg)
}

//根据不同服务需要，执行不同的初始化操作,注意不同初始化操作执行的先后依赖顺序
func serverInit(envName string) {
	dm.InitConf(envName)
}
