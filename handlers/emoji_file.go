package handlers

import (
	"OwlGramServer/telegram/emoji"
	"github.com/valyala/fasthttp"
)

func EmojiFile(ctx *fasthttp.RequestCtx, identifier string, clientEmoji *emoji.Context) {
	for _, pack := range clientEmoji.EmojiPacks {
		if pack.Identifier == identifier {
			ctx.SetContentType("application/zip")
			ctx.SetStatusCode(fasthttp.StatusOK)
			ctx.SetBody(pack.Emojies)
			ctx.SetConnectionClose()
			return
		}
	}
	NotFound(ctx)
}
