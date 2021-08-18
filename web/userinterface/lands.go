package ui

import (
	"web/problemdomain"

	"github.com/fasthttp/router"
)

const Prefix = "/vesta/demo/"

func Init(router *router.Router) {

	//location和执行函数进行绑定
	router.GET(Prefix+"location", pd.Demo)
	router.POST(Prefix+"location", pd.Demo)
}
