package ui

import (
	"github.com/qxyang2015/accumulation/web/problemdomain"

	"github.com/fasthttp/router"
)

const Prefix = "/web/"

func Init(router *router.Router) {
	//location和执行函数进行绑定
	router.GET(Prefix+"demo", pd.Demo)
	router.POST(Prefix+"formdata", pd.FormdataDemo)
}
