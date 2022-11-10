package handlers

import (
	"OwlGramServer/crowdin"
	"OwlGramServer/crowdin/types"
	"encoding/json"
	"github.com/valyala/fasthttp"
)

func LanguageVersion(ctx *fasthttp.RequestCtx, crowdinClient *crowdin.Context) {
	lang := ctx.Request.URI().QueryArgs().Peek("lang")
	if lang != nil {
		if len(crowdinClient.DownloadedLanguages) <= 1 {
			Forbidden(ctx)
			return
		}
		ctx.SetContentType("application/json")
		ctx.SetStatusCode(fasthttp.StatusOK)
		result := crowdinClient.GetLanguageMD5(string(lang))
		if result != nil && len(string(lang)) >= 2 {
			res, _ := json.Marshal(types.LanguageVersion{
				MD5: *result,
			})
			ctx.SetBody(res)
		} else {
			res, _ := json.Marshal(types.ResultError{
				StatusError: "Language pack not found!",
			})
			ctx.SetBody(res)
		}
		ctx.SetConnectionClose()
		return
	}
	Forbidden(ctx)
}
