package pd

import (
	"fmt"
	"github.com/qxyang2015/accumulation/tools/http_tools"
	"github.com/valyala/fasthttp"
)

type BaseResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func FormdataDemo(ctx *fasthttp.RequestCtx) {
	response := BaseResponse{
		Code: 0,
		Msg:  "OK",
	}

	defer func() {
		http_tools.HttpResponse(ctx, response)
	}()
	reader, err := ctx.MultipartForm()
	if err != nil {
		response.Code = -1
		response.Msg = fmt.Sprintf("解析Request错误[%v]", err)
		return
	}
	values := reader.Value
	for key, value := range values {
		fmt.Println(key, value)
	}
	files := reader.File
	for key, file := range files {
		fmt.Println(key, file[0].Filename, file[0].Size)
	}
}
