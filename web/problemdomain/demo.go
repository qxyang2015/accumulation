package pd

import (
	"github.com/qxyang2015/accumulation/tools/http_tools"
	"github.com/valyala/fasthttp"
)

func Demo(ctx *fasthttp.RequestCtx) {
	http_tools.HttpResponse(ctx, "{\"A\":\"123\"}")

}
