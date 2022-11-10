package tg_bot

import (
	"OwlGramServer/consts"
	updates "OwlGramServer/updates/types"
	"OwlGramServer/utilities"
	"github.com/Squirrel-Network/gobotapi"
	"github.com/Squirrel-Network/gobotapi/methods"
	"github.com/Squirrel-Network/gobotapi/types"
	"os"
	"strconv"
	"strings"
)

func (c *Context) handleCallbackQuery(client *gobotapi.Client, callbackQuery types.CallbackQuery) {
	switch callbackQuery.Data {
	case "delete_message":
		_, _ = client.Invoke(&methods.DeleteMessage{
			ChatID:    callbackQuery.Message.Chat.ID,
			MessageID: callbackQuery.Message.MessageID,
		})
		break
	case "recover_message":
		if c.CurrentStatus != 0 {
			_, _ = client.Invoke(&methods.AnswerCallbackQuery{
				CallbackQueryID: callbackQuery.ID,
			})
			c.CacheMessage = *callbackQuery.Message
			c.CompilerClient.MessageID = callbackQuery.Message.MessageID
			c.StatusHandler(c.CurrentStatus)
		} else {
			_, _ = client.Invoke(&methods.AnswerCallbackQuery{
				CallbackQueryID: callbackQuery.ID,
				Text:            "⚠️ Nessun backup trovato",
			})
		}
		break
	default:
		releaseType := strings.Split(callbackQuery.Data, ":")
		if c.CompilerClient.MessageID == callbackQuery.Message.MessageID && c.CompilerClient.IsRunning {
			if len(releaseType) < 2 {
				c.StatusHandler(utilities.BundlesSelectReleaseType)
			} else if releaseType[1] == "cancel" {
				c.StatusHandler(utilities.BundlesCompiled)
			} else if releaseType[1] == "changelogs" {
				_, _ = client.Invoke(&methods.AnswerCallbackQuery{
					CallbackQueryID: callbackQuery.ID,
				})
				c.SelectedLanguage = releaseType[2]
				c.StatusHandler(utilities.NeededChangelogs)
			} else if releaseType[1] == "back_to_changelogs" {
				c.StatusHandler(utilities.NeededChangelogs)
			} else if releaseType[1] == "edit_text" {
				_, _ = client.Invoke(&methods.AnswerCallbackQuery{
					CallbackQueryID: callbackQuery.ID,
				})
				if len(releaseType) < 3 {
					c.StatusHandler(utilities.EditingChangelogs)
				} else if releaseType[2] == "image" {
					c.StatusHandler(utilities.NeededImage)
				} else if releaseType[2] == "desc" {
					c.StatusHandler(utilities.NeededDesc)
				} else if releaseType[2] == "link" {
					c.StatusHandler(utilities.NeededLink)
				} else if releaseType[2] == "note" {
					c.StatusHandler(utilities.NeededNotes)
				}
			} else if releaseType[1] == "send_image" {
				_, _ = client.Invoke(&methods.AnswerCallbackQuery{
					CallbackQueryID: callbackQuery.ID,
				})
				content, _ := os.ReadFile(c.BannerTmp)
				_, _ = client.Invoke(&methods.SendPhoto{
					ChatID:           c.CacheMessage.Chat.ID,
					ReplyToMessageID: c.CacheMessage.MessageID,
					Photo: types.InputBytes{
						Name: "photo.png",
						Data: content,
					},
					ReplyMarkup: &types.InlineKeyboardMarkup{
						InlineKeyboard: [][]types.InlineKeyboardButton{
							{
								{
									Text:         "🗑 Chiudi",
									CallbackData: "delete_message",
								},
							},
						},
					},
				})
			} else if releaseType[1] == "confirm" {
				_, _ = client.Invoke(&methods.AnswerCallbackQuery{
					CallbackQueryID: callbackQuery.ID,
				})
				c.StatusHandler(utilities.ConfirmChanges)
			} else if releaseType[1] == "confirm_changes" {
				updateDescription := c.UpdateClient.GetChangelogUpdate(strconv.Itoa(c.CompilerClient.Bundles[0].VersionCode), "en", true)
				if updateDescription != nil {
					c.StoreClient.Run(c.CompilerClient.Bundles, c.ReleaseType, c.BannerTmp, c.HrefTmp, c.LocalizationsTmp, *c.UpdateClient.UpdatesDescriptor, c.StatusHandler)
					_, _ = client.Invoke(&methods.AnswerCallbackQuery{
						CallbackQueryID: callbackQuery.ID,
					})
					c.CompilerClient.IsRunning = false
					c.CurrentStatus = 0
					c.StatusHandler(utilities.BundlesSentToStores)
					c.StatusHandler(utilities.ChangesConfirmed)
				} else {
					_, _ = client.Invoke(&methods.AnswerCallbackQuery{
						CallbackQueryID: callbackQuery.ID,
						Text:            "⚠️ Changelogs non trovati",
					})
				}
			} else if releaseType[1] == "cancel_release" {
				_, _ = client.Invoke(&methods.AnswerCallbackQuery{
					CallbackQueryID: callbackQuery.ID,
				})
				c.StatusHandler(utilities.BundlesCompiled)
			} else if releaseType[1] == "base" {
				c.ReleaseBase = releaseType[2]
				var updateInfo *updates.UpdateInfo
				if c.ReleaseType == "stable" {
					updateInfo = c.UpdateClient.UpdatesDescriptor.Updates.Stable
				} else {
					updateInfo = c.UpdateClient.UpdatesDescriptor.Updates.Beta
				}
				c.HrefTmp = updateInfo.Href
				c.BannerTmp = strings.ReplaceAll(updateInfo.Banner, consts.OwlGramFilesServer, consts.ServerFilesFolder)
				c.LocalizationsTmp = c.UpdateClient.UpdatesDescriptor.Localizations
				c.SelectedLanguage = "en"
				c.StatusHandler(utilities.NeededChangelogs)
			} else if releaseType[1] == "delete" {
				_, _ = client.Invoke(&methods.DeleteMessage{
					ChatID:    callbackQuery.Message.Chat.ID,
					MessageID: callbackQuery.Message.MessageID,
				})
				c.CurrentStatus = 0
				c.CompilerClient.IsRunning = false
				_, _ = client.Invoke(&methods.SendMessage{
					ChatID:    callbackQuery.Message.Chat.ID,
					Text:      "<i><b>🗑 Rilascio annullato</b></i>",
					ParseMode: "HTML",
				})
			} else {
				c.ReleaseType = releaseType[1]
				c.StatusHandler(utilities.SelectBaseBundle)
			}
		} else {
			_, _ = client.Invoke(&methods.AnswerCallbackQuery{
				CallbackQueryID: callbackQuery.ID,
				Text:            "⚠️ Operazione scaduta",
			})
		}
		break
	}
}
