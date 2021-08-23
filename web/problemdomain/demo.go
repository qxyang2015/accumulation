package pd

import (
	"fmt"
	"github.com/qxyang2015/accumulation/tools/http_tools"
	"github.com/valyala/fasthttp"
)

type DemoRequest struct {
	Id string `json:"id"`
}

func Demo(ctx *fasthttp.RequestCtx) {
	fmt.Println("request:", string(ctx.Request.Body()))
	http_tools.HttpResponse(ctx, string(ctx.Request.Body()))
}
