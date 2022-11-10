package handlers

import (
	"OwlGramServer/updates"
	"OwlGramServer/updates/types"
	"encoding/json"
	"github.com/valyala/fasthttp"
)

func Changelogs(ctx *fasthttp.RequestCtx, updatesClient *updates.Context) {
	lang := ctx.Request.URI().QueryArgs().Peek("lang")
	version := ctx.Request.URI().QueryArgs().Peek("version")
	if lang != nil && version != nil && updatesClient.UpdatesDescriptor != nil {
		ctx.SetContentType("application/json")
		ctx.SetStatusCode(fasthttp.StatusOK)
		res, _ := json.Marshal(types.Changelogs{
			Changelogs: updatesClient.GetChangelogUpdate(string(version), string(lang), false),
		})
		ctx.SetBody(res)
		ctx.SetConnectionClose()
		return
	}
	Forbidden(ctx)
}
