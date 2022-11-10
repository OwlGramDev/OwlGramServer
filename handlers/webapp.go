package handlers

import (
	"OwlGramServer/consts"
	"OwlGramServer/crowdin"
	telegram "OwlGramServer/tg_bot"
	"OwlGramServer/tg_bot/types"
	"OwlGramServer/utilities"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/Squirrel-Network/gobotapi"
	"github.com/Squirrel-Network/gobotapi/methods"
	"github.com/valyala/fasthttp"
	"golang.org/x/exp/utf8string"
	"golang.org/x/net/html"
	"io"
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
		r, _ := os.ReadFile(path.Join(consts.WebAppFiles, "webapp.html"))
		r = bytes.ReplaceAll(r, []byte("%TIME%"), []byte(strconv.Itoa(int(time.Now().Unix()))))
		node, _ := html.Parse(bytes.NewBuffer(r))
		bodyElement := utilities.GetHtmlElementByTag(utilities.GetHtmlElementByTag(node, "html"), "body")
		approvedStrings := crowdinClient.GetAllApprovedStrings()

		for _, nonApprovedString := range approvedStrings {
			result, _ := html.Parse(bytes.NewBuffer([]byte(fmt.Sprintf("<div class=\"section_descriptor\" style=\"display: none;opacity: 0;\">%s</div>", utilities.GetAndroidLangName(nonApprovedString.LanguageCode)))))
			bodyElement.AppendChild(result)
			for _, stringToApprove := range nonApprovedString.Vars {
				var buf bytes.Buffer
				w := io.Writer(&buf)
				_ = html.Render(w, bodyElement.FirstChild.NextSibling)
				translatorName := stringToApprove.TranslatorName
				s := utf8string.NewString(translatorName)
				if s.RuneCount() == 0 {
					continue
				}
				htmlNew := fmt.Sprintf(
					buf.String(),
					strings.ToUpper(s.Slice(0, 1)),
					stringToApprove.Lang,
					fmt.Sprintf("New Translation by %s", stringToApprove.TranslatorName),
					stringToApprove.Name,
					html.EscapeString(stringToApprove.Text),
					stringToApprove.TranslationId,
					stringToApprove.TranslationId,
				)
				result, _ = html.Parse(bytes.NewBuffer([]byte(htmlNew)))
				bodyElement.AppendChild(result)
			}
		}
		bodyElement.RemoveChild(bodyElement.FirstChild.NextSibling)
		var buf bytes.Buffer
		w := io.Writer(&buf)
		_ = html.Render(w, node)
		ctx.SetBody(buf.Bytes())
	}
	ctx.SetConnectionClose()
}
