package main

import (
	"OwlGramServer/consts"
	"OwlGramServer/github_bot"
	"OwlGramServer/handlers"
	"github.com/valyala/fasthttp"
)

func handler(ctx *fasthttp.RequestCtx) {
	ip := string(ctx.Request.Header.Peek("X-Forwarded-For"))
	if consts.IsDebug && ip != consts.SshIP {
		handlers.Forbidden(ctx)
		return
	}
	if ctx.IsGet() {
		switch string(ctx.Path()) {
		case "/get_changelogs":
			handlers.Changelogs(ctx, updatesClient)
		case "/version":
			handlers.Updates(ctx, updatesClient)
		case "/dc_status":
			handlers.DcStatus(ctx, tgChecker)
		case "/language_pack":
			handlers.LanguagePack(ctx, crowdinClient)
		case "/language_version":
			handlers.LanguageVersion(ctx, crowdinClient)
		case "/webapp":
			handlers.WebApp(ctx, crowdinClient, botClient.TelegramClient)
		default:
			handlers.NotFound(ctx)
		}
	} else if ctx.IsPost() {
		switch string(ctx.Path()) {
		case "/saveDcData":
			handlers.SaveDCData(ctx, tgChecker)
		case "/notify_upload":
			handlers.NotifyUpload(ctx, botClient)
		case "/notify_github":
			github_bot.SendPushEvent(ctx)
		default:
			handlers.NotFound(ctx)
			break
		}
	} else {
		handlers.Forbidden(ctx)
	}
}
