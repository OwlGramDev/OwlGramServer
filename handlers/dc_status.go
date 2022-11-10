package handlers

import (
	"OwlGramServer/tg_checker"
	"OwlGramServer/tg_checker/types"
	"encoding/json"
	"github.com/valyala/fasthttp"
	"time"
)

func DcStatus(ctx *fasthttp.RequestCtx, tgChecker *tg_checker.Context) {
	ctx.SetContentType("application/json")
	ctx.SetStatusCode(fasthttp.StatusOK)
	ctx.Response.Header.Set("Access-Control-Allow-Origin", "https://status.owlgram.org")
	if tgChecker == nil {
		Forbidden(ctx)
		return
	}
	updateInTime := int8(60 - time.Now().Second())
	if updateInTime == 0 {
		tgChecker.IsRefreshing = true
	}
	if tgChecker.IsRefreshing {
		updateInTime = 0
	}
	res, _ := json.Marshal(types.JSONResult{
		Status:        tgChecker.StatusDC,
		LastRefresh:   tgChecker.LastRefresh,
		RefreshInTime: updateInTime,
		IsRefreshing:  tgChecker.IsRefreshing,
	})
	ctx.SetBody(res)
	ctx.SetConnectionClose()
}
