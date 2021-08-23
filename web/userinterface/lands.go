package ui

import (
	"github.com/fasthttp/router"
	"github.com/qxyang2015/accumulation/web/problemdomain"
	"github.com/valyala/fasthttp"
)

const Prefix = "/web/"

func Init(router *router.Router) {
	//location和执行函数进行绑定
	r := NewRouter()
	r.Use(TimeMiddleWare)
	r.Add(fasthttp.MethodPost, Prefix+"demo", pd.Demo)

	//router.POST(Prefix+"formdata", pd.FormdataDemo)
	for method, urlMap := range r.routerMap {
		for url, handler := range urlMap {
			router.Handle(method, url, handler)
		}
	}
}
