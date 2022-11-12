package handlers

import (
	"OwlGramServer/consts"
	"OwlGramServer/crowdin"
	telegram "OwlGramServer/tg_bot"
	"OwlGramServer/tg_bot/types"
	"OwlGramServer/utilities"
	"encoding/json"
	"github.com/Squirrel-Network/gobotapi"
	"github.com/Squirrel-Network/gobotapi/methods"
	"github.com/flosch/pongo2"
	"github.com/valyala/fasthttp"
	"golang.org/x/exp/utf8string"
	"golang.org/x/net/html"
	"os"
	"path"
	"strconv"
	"strings"
	"time"
)

func WebApp(ctx *fasthttp.RequestCtx, crowdinClient *crowdin.Context, telegramClient *gobotapi.PollingClient) {
	action := ctx.Request.URI().QueryArgs().Peek("action")
	ctx.SetStatusCode(fasthttp.StatusOK)
	if action != nil {
		secret := ctx.Request.URI().QueryArgs().Peek("secret")
		if secret == nil {
			Forbidden(ctx)
			return
		}
		if !telegram.CheckTelegramAuthorization(secret) {
			Forbidden(ctx)
			return
		}
		ctx.SetContentType("application/json")
		switch string(action) {
		case "getAdmins":
			marshal, err := json.Marshal(consts.AuthorizedUsers)
			if err != nil {
				return
			}
			ctx.SetBody(marshal)
			break
		case "confirmChanges":
			data := ctx.Request.URI().QueryArgs().Peek("data")
			if data == nil {
				return
			}
			var result types.WebAppInit
			_ = json.Unmarshal(secret, &result)
			authDate, _ := strconv.Atoi(result.AuthDate)
			if !utilities.Contains(consts.AuthorizedUsers, result.User.ID) || time.Now().Unix()-int64(authDate) > consts.WebAppAuthTimeout {
				return
			}
			go func() {
				var dataStruct struct {
					Delete  []int `json:"delete"`
					Approve []int `json:"approve"`
				}
				_ = json.Unmarshal(data, &dataStruct)
				for _, id := range dataStruct.Approve {
					go crowdinClient.ApproveString(id)
				}
				for _, id := range dataStruct.Delete {
					go crowdinClient.DeleteString(id)
				}
				_, _ = telegramClient.Invoke(&methods.SendMessage{
					ChatID:    result.User.ID,
					Text:      "<b>✅ Changes applied</b>",
					ParseMode: "HTML",
				})
			}()
			ctx.SetBodyString("{\"status\":\"ok\"}")
			break
		}
	} else {
		ctx.SetContentType("text/html; charset=UTF-8")
		file, _ := os.ReadFile(path.Join(consts.WebAppFiles, "webapp.html"))
		tmp, err := pongo2.FromBytes(file)
		if err != nil {
			Forbidden(ctx)
			return
		}
		resBuild, err := tmp.Execute(pongo2.Context{
			"currentTime":           time.Now().Unix(),
			"getAllApprovedStrings": crowdinClient.GetAllApprovedStrings,
			"getShortName": func(name string) string {
				s := utf8string.NewString(name)
				return strings.ToUpper(s.Slice(0, 1))
			},
			"getAndroidLangName": utilities.GetAndroidLangName,
			"escapeString":       html.EscapeString,
		})
		if err != nil {
			Forbidden(ctx)
			return
		}
		ctx.SetBody([]byte(resBuild))
	}
	ctx.SetConnectionClose()
}
