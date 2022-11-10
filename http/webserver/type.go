package webserver

import "github.com/valyala/fasthttp"

type Context struct {
	handler func(ctx *fasthttp.RequestCtx)
}
