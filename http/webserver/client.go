package webserver

import "github.com/valyala/fasthttp"

func Client(handler func(ctx *fasthttp.RequestCtx)) *Context {
	client := &Context{
		handler: handler,
	}
	return client
}
