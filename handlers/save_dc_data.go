package handlers

import (
	"OwlGramServer/consts"
	"OwlGramServer/tg_checker"
	"OwlGramServer/tg_checker/types"
	"encoding/json"
	"github.com/valyala/fasthttp"
	"time"
)

func SaveDCData(ctx *fasthttp.RequestCtx, tgChecker *tg_checker.Context) {
	secretKey := ctx.Request.URI().QueryArgs().Peek("secretKey")
	if secretKey != nil {
		if string(secretKey) == consts.SecretDCKey {
			var listStatus []types.DCStatus
			err := json.Unmarshal(ctx.PostBody(), &listStatus)
			if err == nil {
				ctx.SetStatusCode(fasthttp.StatusOK)
				tgChecker.StatusDC = listStatus
				tgChecker.LastRefresh = time.Now().Unix()
				tgChecker.IsRefreshing = false
			}
		}
	}
	Forbidden(ctx)
}
