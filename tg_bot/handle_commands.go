package tg_bot

import (
	"OwlGramServer/consts"
	"OwlGramServer/utilities"
	"fmt"
	"github.com/Squirrel-Network/gobotapi"
	"github.com/Squirrel-Network/gobotapi/filters"
	"github.com/Squirrel-Network/gobotapi/methods"
	"github.com/Squirrel-Network/gobotapi/types"
)

func (c *Context) handleCommands() {
	c.TelegramClient.OnAnyMessageEvent(filters.Filter(func(client *gobotapi.Client, message types.Message) {
		if message.Chat.ID != message.From.ID {
			return
		}
		_, _ = client.Invoke(&methods.SendMessage{
			ChatID:           message.Chat.ID,
			Text:             "<b>Welcome!</b>\n\nThis bot is used to translate manage Crowdin translations.",
			ReplyToMessageID: message.MessageID,
			ParseMode:        "HTML",
			ReplyMarkup: &types.InlineKeyboardMarkup{
				InlineKeyboard: [][]types.InlineKeyboardButton{
					{
						{
							Text: "Translation Platform",
							WebApp: &types.WebAppInfo{
								URL: consts.WebAppLink,
							},
						},
					},
				},
			},
		})
	}, filters.Command("start", consts.AliasSupported...)))
	c.TelegramClient.OnAnyMessageEvent(filters.Filter(func(client *gobotapi.Client, message types.Message) {
		if c.CompilerClient.IsRunning {
			_, _ = client.Invoke(&methods.SendMessage{
				ChatID:           message.Chat.ID,
				Text:             "<i>\u274C Un altro processo è già in esecuzione!</i>",
				ParseMode:        "HTML",
				ReplyToMessageID: message.MessageID,
				ReplyMarkup: &types.InlineKeyboardMarkup{
					InlineKeyboard: [][]types.InlineKeyboardButton{
						{
							{
								Text:         "♻️ Recupera Messaggio",
								CallbackData: "recover_message",
							},
						},
						{
							{
								Text:         "🗑 Chiudi",
								CallbackData: "delete_message",
							},
						},
					},
				},
			})
		} else {
			c.CompilerClient.IsRunning = true
			c.CurrentStatus = 0
			resultMessage, err := client.Invoke(&methods.SendMessage{
				ChatID:           message.Chat.ID,
				Text:             "<i>\U0001F4E6 Estrazione in corso...</i>",
				ReplyToMessageID: message.MessageID,
				ParseMode:        "HTML",
			})
			if err != nil {
				return
			}
			c.CacheMessage = resultMessage.Result.(types.Message)
			c.CompilerClient.Run(resultMessage.Result.(types.Message).MessageID, c.StatusHandler)
		}
	}, filters.And(filters.UserID(consts.AuthorizedUsers...), filters.Command("publish", consts.AliasSupported...))))
	c.TelegramClient.OnAnyMessageEvent(filters.Filter(func(client *gobotapi.Client, message types.Message) {
		if !c.SendingApks && c.CurrentStatus >= utilities.NeededChangelogs && c.CompilerClient.IsRunning {
			c.SendingApks = true
			resultMessage, _ := client.Invoke(&methods.SendMessage{
				ChatID:           message.Chat.ID,
				Text:             "<i>📤 Invio apk corso...</i>",
				ReplyToMessageID: message.MessageID,
				ParseMode:        "HTML",
			})
			err := c.SendApks(message.Chat.ID)
			if err != nil {
				_, _ = client.Invoke(&methods.SendMessage{
					ChatID:           message.Chat.ID,
					Text:             fmt.Sprintf("<i>📤 Errore durante l'invio degli apk</i>\n\n<b>Errore:</b> <code>%s</code>" + err.Error()),
					ReplyToMessageID: message.MessageID,
					ParseMode:        "HTML",
				})
			}
			_, _ = client.Invoke(&methods.DeleteMessage{
				ChatID:    message.Chat.ID,
				MessageID: resultMessage.Result.(types.Message).MessageID,
			})
			c.SendingApks = false
		}
	}, filters.And(filters.UserID(consts.AuthorizedUsers...), filters.Command("apk", consts.AliasSupported...))))
}
