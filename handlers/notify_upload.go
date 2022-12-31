package handlers

import (
	"OwlGramServer/consts"
	"OwlGramServer/telegram/bot"
	"github.com/Squirrel-Network/gobotapi/methods"
	"github.com/valyala/fasthttp"
)

func NotifyUpload(ctx *fasthttp.RequestCtx, bot *bot.Context) {
	token := ctx.Request.Header.Peek("Token")
	if token == nil {
		Forbidden(ctx)
		return
	}
	if string(token) != consts.PublisherToken {
		Forbidden(ctx)
		return
	}
	ctx.SetContentType("application/json")
	ctx.SetStatusCode(fasthttp.StatusOK)
	_, _ = ctx.WriteString("ok")
	text := "<b>🎉 Novità</b>\n"
	text += "Lo sviluppatore di OwlGram ha pubblicato nuovi bundles!\n\n"
	text += "<i>Esegui /publish per avviare la pubblicazione.</i>\n"
	_, _ = bot.TelegramClient.Invoke(&methods.SendMessage{
		ChatID:    consts.StaffGroupID,
		Text:      text,
		ParseMode: "HTML",
	})
	ctx.SetConnectionClose()
}
