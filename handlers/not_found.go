package handlers

import (
	"OwlGramServer/consts"
	"github.com/valyala/fasthttp"
	"os"
)

func NotFound(ctx *fasthttp.RequestCtx) {
	ctx.SetContentType("text/html")
	ctx.SetStatusCode(fasthttp.StatusNotFound)
	r, err := os.ReadFile(consts.NotFoundFile)
	if err == nil {
		ctx.SetBody(r)
	} else {
		ctx.SetBodyString("404 Not Found")
	}
	ctx.SetConnectionClose()
}
