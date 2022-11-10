package handlers

import (
	"OwlGramServer/consts"
	"github.com/valyala/fasthttp"
	"os"
)

func Forbidden(ctx *fasthttp.RequestCtx) {
	ctx.SetContentType("text/html")
	ctx.SetStatusCode(fasthttp.StatusForbidden)
	r, err := os.ReadFile(consts.ForbiddenFile)
	if err == nil {
		ctx.SetBody(r)
	} else {
		ctx.SetBodyString("<h1>Forbidden</h1>")
	}
	ctx.SetConnectionClose()
}
