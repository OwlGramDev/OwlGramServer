package handlers

import (
	"OwlGramServer/emoji"
	"github.com/valyala/fasthttp"
)

func EmojiPreview(ctx *fasthttp.RequestCtx, identifier string, clientEmoji *emoji.Context) {
	for _, pack := range clientEmoji.EmojiPacks {
		if pack.ID == identifier {
			ctx.SetContentType("image/png")
			ctx.SetStatusCode(fasthttp.StatusOK)
			ctx.SetBody(pack.Preview)
			ctx.SetConnectionClose()
			return
		}
	}
	NotFound(ctx)
}
