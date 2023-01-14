package handlers

import (
	"OwlGramServer/emoji"
	"github.com/valyala/fasthttp"
)

func EmojiFile(ctx *fasthttp.RequestCtx, identifier string, clientEmoji *emoji.Context) {
	for _, pack := range clientEmoji.EmojiPacks {
		if pack.ID == identifier {
			ctx.SetContentType("application/zip")
			ctx.SetStatusCode(fasthttp.StatusOK)
			ctx.SetBody(pack.EmojiZip)
			ctx.SetConnectionClose()
			return
		}
	}
	NotFound(ctx)
}
