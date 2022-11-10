package tg_bot

import (
	"OwlGramServer/compiler"
	"OwlGramServer/consts"
	"OwlGramServer/utilities"
	"fmt"
	"github.com/Squirrel-Network/gobotapi"
	"github.com/Squirrel-Network/gobotapi/methods"
	"github.com/Squirrel-Network/gobotapi/types"
	"regexp"
	"strconv"
)

func (c *Context) handleMessage(client *gobotapi.Client, message types.Message) {
	sentResponse := c.CompilerClient.IsRunning
	sentResponse = sentResponse && message.ReplyToMessage != nil
	sentResponse = sentResponse && message.ReplyToMessage.MessageID == c.CacheMessage.MessageID
	var actionRun func()
	switch c.CurrentStatus {
	case utilities.NeededImage:
		if message.Document != nil && (message.Document.MimeType == "image/png" || message.Document.MimeType == "image/jpeg" || message.Document.MimeType == "image/jpg") {
			actionRun = func() {
				outputPath := fmt.Sprintf("%stmp.png", consts.BotImagesCache)
				_ = client.DownloadMedia(message, outputPath, nil)
				c.BannerTmp = outputPath
			}
		}
		break
	case utilities.NeededDesc:
		if message.Text != "" {
			actionRun = func() {
				c.LocalizationsTmp[c.SelectedLanguage][fmt.Sprintf("desc2_%s", c.ReleaseType)] = message.Text
			}
		}
		break
	case utilities.NeededNotes:
		if message.Text != "" {
			actionRun = func() {
				c.LocalizationsTmp[c.SelectedLanguage][fmt.Sprintf("note_%s", c.ReleaseType)] = message.Text
			}
		}
		break
	case utilities.NeededLink:
		if message.Text != "" {
			actionRun = func() {
				c.HrefTmp = message.Text
			}
		}
		break
	}
	if message.Document != nil && (message.Document.MimeType == "application/octet-stream" || message.Document.MimeType == "message/rfc822") {
		compile, _ := regexp.Compile("\\((.*?)\\)")
		matches := compile.FindAllStringSubmatch(message.Caption, -1)
		var versionCheck int
		if len(matches) > 0 {
			resSend, _ := client.Invoke(&methods.SendMessage{
				ChatID:           message.Chat.ID,
				Text:             "<i>🔄 Deoffuscamento in corso...</i>",
				ParseMode:        "HTML",
				ReplyToMessageID: message.MessageID,
			})
			versionCheck, _ = strconv.Atoi(matches[0][1])
			result, _ := client.DownloadBytes(message.Document.FileID, nil)
			res, err := compiler.DeObfuscateLogs(string(result), versionCheck)
			if len(res) > 0 {
				_, _ = client.Invoke(&methods.DeleteMessage{
					ChatID:    message.Chat.ID,
					MessageID: resSend.Result.(types.Message).MessageID,
				})
				_, _ = client.Invoke(&methods.SendDocument{
					ChatID: message.Chat.ID,
					Document: types.InputBytes{
						Name: "Logcat.txt",
						Data: res,
					},
					ReplyToMessageID: message.MessageID,
				})
			} else {
				_, _ = client.Invoke(&methods.DeleteMessage{
					ChatID:    message.Chat.ID,
					MessageID: resSend.Result.(types.Message).MessageID,
				})
				_, _ = client.Invoke(&methods.SendMessage{
					ChatID:           message.Chat.ID,
					Text:             fmt.Sprintf("<i>❌ Deoffuscamento fallito!</i>\n<b>Errore:</b> <code>%s</code>", err.Error()),
					ParseMode:        "HTML",
					ReplyToMessageID: message.MessageID,
				})
			}
		}
	}
	if sentResponse && actionRun != nil {
		res, _ := client.Invoke(&methods.SendMessage{
			ChatID:           message.Chat.ID,
			Text:             "<i>🔄 Caricamento in corso...</i>",
			ParseMode:        "HTML",
			ReplyToMessageID: message.MessageID,
		})
		actionRun()
		_, _ = client.Invoke(&methods.DeleteMessage{
			ChatID:    c.CacheMessage.Chat.ID,
			MessageID: c.CacheMessage.MessageID,
		})
		c.CacheMessage = res.Result.(types.Message)
		c.CompilerClient.MessageID = res.Result.(types.Message).MessageID
		c.StatusHandler(utilities.EditedChangelogs)
		return
	}
}
