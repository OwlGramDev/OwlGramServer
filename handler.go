package main

import (
	"OwlGramServer/consts"
	"OwlGramServer/github_bot"
	"OwlGramServer/handlers"
	"github.com/valyala/fasthttp"
	"strings"
)

func handler(ctx *fasthttp.RequestCtx) {
	ip := string(ctx.Request.Header.Peek("X-Forwarded-For"))
	if consts.IsDebug && ip != consts.SshIP {
		handlers.Forbidden(ctx)
		return
	}
	requestedPath := string(ctx.Path())
	if ctx.IsGet() {
		switch requestedPath {
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
		case "/emoji_packs":
			handlers.EmojiPacks(ctx, emojiClient)
		default:
			if len(requestedPath) > 14 && strings.HasPrefix(requestedPath, "/previews/") &&
				strings.HasSuffix(requestedPath, ".png") {
				requestedPath = requestedPath[10 : len(requestedPath)-4]
				handlers.EmojiPreview(ctx, requestedPath, emojiClient)
				return
			} else if len(requestedPath) > 10 && strings.HasPrefix(requestedPath, "/packs/") &&
				strings.HasSuffix(requestedPath, ".zip") {
				requestedPath = requestedPath[7 : len(requestedPath)-4]
				handlers.EmojiFile(ctx, requestedPath, emojiClient)
				return
			}
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
