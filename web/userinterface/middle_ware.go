package ui

import (
	"fmt"
	"github.com/valyala/fasthttp"
	"time"
)

//请求中间件的一种实现方式
type MiddleWareFunc func(fasthttp.RequestHandler) fasthttp.RequestHandler

type MiddleRouter struct {
	middleWareChain []MiddleWareFunc
	routerMap       map[string]map[string]fasthttp.RequestHandler
}

func NewRouter() *MiddleRouter {
	return &MiddleRouter{
		routerMap: make(map[string]map[string]fasthttp.RequestHandler),
	}
}

func (mr *MiddleRouter) Use(mf MiddleWareFunc) {
	mr.middleWareChain = append(mr.middleWareChain, mf)
}

func (mr *MiddleRouter) Add(method, url string, h fasthttp.RequestHandler) {
	mergeHandler := h
	for i := len(mr.middleWareChain) - 1; i >= 0; i-- {
		mergeHandler = mr.middleWareChain[i](mergeHandler)
	}
	if _, ok := mr.routerMap[method]; !ok {
		mr.routerMap[method] = make(map[string]fasthttp.RequestHandler)
	}
	mr.routerMap[method][url] = mergeHandler
}

func TimeMiddleWare(next fasthttp.RequestHandler) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		timeStart := time.Now()
		next(ctx)
		timeElapsed := time.Since(timeStart)
		fmt.Println(timeElapsed)
	}
}
