package ui

import (
	"github.com/qxyang2015/accumulation/web/problemdomain"
	"github.com/valyala/fasthttp"

	"github.com/fasthttp/router"
)

const Prefix = "/web/"

func Init(router *router.Router) {
	//location和执行函数进行绑定
	router.POST(Prefix+"demo", pd.Demo)
	router.POST(Prefix+"formdata", pd.FormdataDemo)
}

//请求中间件的一种实现方式
type MiddleWareFunc func(fasthttp.RequestHandler) fasthttp.RequestHandler

type MiddleRouter struct {
	middleWareChain []MiddleWareFunc
	routerMap       map[string]fasthttp.RequestHandler
}

func NewRouter() *MiddleRouter {
	return &MiddleRouter{
		routerMap: make(map[string]fasthttp.RequestHandler),
	}
}

func (mr *MiddleRouter) Use(mf MiddleWareFunc) {
	mr.middleWareChain = append(mr.middleWareChain, mf)
}

func (mr *MiddleRouter) Add(url string, h fasthttp.RequestHandler) {
	mergeHandler := h
	for i := len(mr.middleWareChain) - 1; i >= 0; i++ {
		mergeHandler = mr.middleWareChain[i](mergeHandler)
	}
	mr.routerMap[url] = mergeHandler
}
