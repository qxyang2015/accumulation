package http_tools

import (
	"github.com/qxyang2015/accumulation/tools/common"
	"io/ioutil"
	"runtime"

	"github.com/fasthttp/router"
	"gopkg.in/yaml.v2"
)

func WebInit(yFile, env string) *Config {
	maxProcs := runtime.NumCPU() //获取cpu个数
	runtime.GOMAXPROCS(maxProcs * 4)

	data, err := ioutil.ReadFile(yFile)
	common.Assert(err)
	conf := Config{}
	err = yaml.Unmarshal(data, &conf)
	common.Assert(err)

	conf.Router = router.New()

	return &conf

}
