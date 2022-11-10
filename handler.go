package main

import (
	"OwlGramServer/consts"
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
			break
		case "/version":
			handlers.Updates(ctx, updatesClient)
			break
		case "/dc_status":
			handlers.DcStatus(ctx, tgChecker)
			break
		case "/language_pack":
			handlers.LanguagePack(ctx, crowdinClient)
			break
		case "/language_version":
			handlers.LanguageVersion(ctx, crowdinClient)
			break
		case "/webapp":
			handlers.WebApp(ctx, crowdinClient, botClient.TelegramClient)
			break
		default:
			handlers.NotFound(ctx)
			break
		}
	} else if ctx.IsPost() {
		switch string(ctx.Path()) {
		case "/saveDcData":
			handlers.SaveDCData(ctx, tgChecker)
			break
		case "/notify_upload":
			handlers.NotifyUpload(ctx, botClient)
			break
		default:
			handlers.NotFound(ctx)
			break
		}
	} else {
		handlers.Forbidden(ctx)
	}
}
