package pd

import (
	"vesta_mlp/util/http_tools"
	"vesta_mlp/util/logging"

	"github.com/valyala/fasthttp"
)

func Demo(ctx *fasthttp.RequestCtx) {
	log4sys := logging.GetLogger()

	log4sys.Trace("这是一个web demo例子")

	http_tools.HttpResponse(ctx, "{\"A\":\"123\"}")

}
