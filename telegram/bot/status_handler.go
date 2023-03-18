package bot

import (
	"OwlGramServer/compiler"
	"OwlGramServer/consts"
	"OwlGramServer/utilities"
	"context"
	"fmt"
	"github.com/Squirrel-Network/gobotapi/methods"
	"github.com/Squirrel-Network/gobotapi/types"
	"github.com/google/go-github/v45/github"
	"regexp"
	"sort"
	"strings"
)

func (c *Context) StatusHandler(status int) {
	c.CurrentStatus = status
	switch status {
	case utilities.NoZipFound:
		_, _ = c.TelegramClient.Invoke(&methods.EditMessageText{
			ChatID:    c.CacheMessage.Chat.ID,
			MessageID: c.CacheMessage.MessageID,
			ParseMode: "HTML",
			Text:      "<i>❌ Non ho trovato nessun file zip.</i>",
		})
		break
	case utilities.UnzipError:
		_, _ = c.TelegramClient.Invoke(&methods.EditMessageText{
			ChatID:    c.CacheMessage.Chat.ID,
			MessageID: c.CacheMessage.MessageID,
			ParseMode: "HTML",
			Text:      "<i>❌ Errore durante l'estrazione del file zip.</i>",
		})
		break
	case utilities.NoBundleFound:
		_, _ = c.TelegramClient.Invoke(&methods.EditMessageText{
			ChatID:    c.CacheMessage.Chat.ID,
			MessageID: c.CacheMessage.MessageID,
			ParseMode: "HTML",
			Text:      "<i>❌ Nessun file bundle trovato!</i>",
		})
		break
	case utilities.ReadingBundlesInfo:
		_, _ = c.TelegramClient.Invoke(&methods.EditMessageText{
			ChatID:    c.CacheMessage.Chat.ID,
			MessageID: c.CacheMessage.MessageID,
			ParseMode: "HTML",
			Text:      "<i>🔄 Leggendo le informazioni dei file bundle...</i>",
		})
		break
	case utilities.BundlesHaveDifferentVersions:
		_, _ = c.TelegramClient.Invoke(&methods.EditMessageText{
			ChatID:    c.CacheMessage.Chat.ID,
			MessageID: c.CacheMessage.MessageID,
			ParseMode: "HTML",
			Text:      "<i>❌ I file bundle hanno versioni diverse!</i>",
		})
		break
	case utilities.BundleCorrupted:
		_, _ = c.TelegramClient.Invoke(&methods.EditMessageText{
			ChatID:    c.CacheMessage.Chat.ID,
			MessageID: c.CacheMessage.MessageID,
			ParseMode: "HTML",
			Text:      "<i>❌ Uno dei file bundle è corrotto!</i>",
		})
		break
	case utilities.BundlesCompiled, utilities.BundlesSelectReleaseType, utilities.SelectBaseBundle:
		var text string
		if status == utilities.BundlesCompiled || status == utilities.BundlesSelectReleaseType {
			text = "<b>\U0001F4A1 OwlGram %s (%d) è pronto!</b>\n\n"
		} else {
			text = "<b>\U0001F4A1 OwlGram %s (%d) è in compilazione.</b>\n\n"
		}
		text = fmt.Sprintf(text, c.CompilerClient.Bundles[0].VersionName, c.CompilerClient.Bundles[0].VersionCode)
		for _, bundle := range c.CompilerClient.Bundles {
			if bundle.IsApk {
				text += fmt.Sprintf("<b>• %s</b>\n", compiler.GetCoolAbi(bundle.AbiName, 2))
			}
		}
		if status == utilities.BundlesCompiled {
			text += "\n<i>Premi il pulsante per inviarlo negli store.</i>"
			_, _ = c.TelegramClient.Invoke(&methods.EditMessageText{
				ChatID:    c.CacheMessage.Chat.ID,
				MessageID: c.CacheMessage.MessageID,
				Text:      text,
				ParseMode: "HTML",
				ReplyMarkup: &types.InlineKeyboardMarkup{
					InlineKeyboard: [][]types.InlineKeyboardButton{
						{
							{
								Text:         "📤 Invia",
								CallbackData: "release",
							},
						},
						{
							{
								Text:         "❌ Annulla",
								CallbackData: "release:delete",
							},
						},
					},
				},
			})
		} else if status == utilities.BundlesSelectReleaseType {
			text += "\n<i>Scegli il tipo di rilascio.</i>"
			_, _ = c.TelegramClient.Invoke(&methods.EditMessageText{
				ChatID:    c.CacheMessage.Chat.ID,
				MessageID: c.CacheMessage.MessageID,
				Text:      text,
				ParseMode: "HTML",
				ReplyMarkup: &types.InlineKeyboardMarkup{
					InlineKeyboard: [][]types.InlineKeyboardButton{
						{
							{
								Text:         "📦 Stabile",
								CallbackData: "release:stable",
							},
							{
								Text:         "🅱️ Beta",
								CallbackData: "release:beta",
							},
						},
						{
							{
								Text:         "❌ Annulla",
								CallbackData: "release:cancel",
							},
						},
					},
				},
			})
		} else {
			ctx := context.Background()
			text += "\n<i>Scegli la versione sui cui è basato questo aggiornamento.</i>"
			client := github.NewClient(nil)
			commits, _, _ := client.Repositories.ListCommits(ctx, "DrkLO", "Telegram", nil)
			var buttons [][]types.InlineKeyboardButton
			max := 4
			for i := 0; i < max; i++ {
				version := *commits[i].GetCommit().Message
				r, _ := regexp.Compile(`\d+(\.\d+)+`)
				versionData := r.FindStringSubmatch(version)
				if len(versionData) == 0 {
					max++
					continue
				}
				version = versionData[0]
				buttons = append(buttons, []types.InlineKeyboardButton{
					{
						Text:         version,
						CallbackData: fmt.Sprintf("release:base:%s", version),
					},
				})
			}
			buttons = append(buttons, []types.InlineKeyboardButton{
				{
					Text:         "🔙 Indietro",
					CallbackData: "cancel_release",
				},
			})
			_, _ = c.TelegramClient.Invoke(&methods.EditMessageText{
				ChatID:    c.CacheMessage.Chat.ID,
				MessageID: c.CacheMessage.MessageID,
				Text:      text,
				ParseMode: "HTML",
				ReplyMarkup: &types.InlineKeyboardMarkup{
					InlineKeyboard: buttons,
				},
			})
		}
		break
	case utilities.SendingToStores, utilities.FailedToSendToStores, utilities.BundlesSentToStores:
		var text string
		if status == utilities.SendingToStores {
			text = "<b>\U0001F4A1 OwlGram %s (%d) è in fase di rilascio sugli store...</b>\n\n"
		} else if status == utilities.FailedToSendToStores {
			text = "<b>\u274C OwlGram %s (%d) non è stato caricato sugli store!</b>\n\n"
		} else {
			text = "<b>\U0001F4A1 OwlGram %s (%d) è stato caricato sugli store!</b>\n\n"
		}
		text = fmt.Sprintf(text, c.CompilerClient.Bundles[0].VersionName, c.CompilerClient.Bundles[0].VersionCode)
		for _, store := range c.StoreClient.StoreList {
			if store.Status == utilities.StatusUploading {
				text += fmt.Sprintf("<b>• %s:</b> <i>%s %s (%v%%)</i>\n", store.Name, utilities.GetEmojiStatus(store.Status), utilities.GetTextStatus(store.Status), store.Progress)
			} else {
				text += fmt.Sprintf("<b>• %s:</b> <i>%s %s</i>\n", store.Name, utilities.GetEmojiStatus(store.Status), utilities.GetTextStatus(store.Status))
			}
		}
		if status == utilities.SendingToStores {
			text += "\n<i>Il processo potrebbe richiedere alcuni minuti...</i>"
		}
		_, _ = c.TelegramClient.Invoke(&methods.EditMessageText{
			ChatID:    c.CacheMessage.Chat.ID,
			MessageID: c.CacheMessage.MessageID,
			Text:      text,
			ParseMode: "HTML",
		})
		if status == utilities.FailedToSendToStores {
			c.CompilerClient.IsRunning = false
		}
		break
	case utilities.NeededChangelogs:
		newLocalization := c.LocalizationsTmp[c.SelectedLanguage]
		text := fmt.Sprintf("<b>💡 OwlGram %s (%d) Changelogs</b>", c.CompilerClient.Bundles[0].VersionName, c.CompilerClient.Bundles[0].VersionCode)
		text += "\n\n"
		text += fmt.Sprintf("<b>• Lingua: %s %s</b>\n\n", utilities.GetEmojiLang(c.SelectedLanguage), utilities.GetLangName(c.SelectedLanguage))
		text += fmt.Sprintf("<b>• Descrizione: </b><i>%s</i>\n\n", newLocalization[fmt.Sprintf("desc2_%s", c.ReleaseType)])
		text += fmt.Sprintf("<b>• Link: </b><i>%s</i>\n\n", c.HrefTmp)
		text += fmt.Sprintf("<b>• Note di rilascio: </b><i>%s</i>", newLocalization[fmt.Sprintf("note_%s", c.ReleaseType)])

		text += "\n\n"
		var buttons []types.InlineKeyboardButton
		localizations := c.UpdateClient.UpdatesDescriptor.Localizations
		keys := make([]string, 0, len(localizations))
		for k := range localizations {
			if k != c.SelectedLanguage {
				keys = append(keys, k)
			}
		}
		sort.Strings(keys)
		for _, k := range keys {
			buttons = append(buttons, types.InlineKeyboardButton{
				Text:         fmt.Sprintf("%s %s", utilities.GetEmojiLang(k), strings.ToUpper(k)),
				CallbackData: fmt.Sprintf("release:changelogs:%s", k),
			})
		}
		var nameCopy string
		if c.ReleaseType == "stable" {
			nameCopy = "Beta"
		} else {
			nameCopy = "Stabile"
		}
		_, _ = c.TelegramClient.Invoke(&methods.EditMessageText{
			ChatID:    c.CacheMessage.Chat.ID,
			MessageID: c.CacheMessage.MessageID,
			Text:      text,
			ParseMode: "HTML",
			ReplyMarkup: &types.InlineKeyboardMarkup{
				InlineKeyboard: [][]types.InlineKeyboardButton{
					buttons,
					{
						{
							Text:         "✏️ Modifica",
							CallbackData: "release:edit_text",
						},
						{
							Text:         fmt.Sprintf("📋 Copia Descrizione (%s)", nameCopy),
							CallbackData: "release:copy_desc",
						},
					},
					{
						{
							Text:         "🖼 Immagine",
							CallbackData: "release:send_image",
						},
					},
					{
						{
							Text:         "❌ Annulla",
							CallbackData: "release:cancel_release",
						},
						{
							Text:         "✅ Conferma",
							CallbackData: "release:confirm",
						},
					},
				},
			},
		})
		break
	case utilities.EditingChangelogs, utilities.EditedChangelogs:
		text := "<i><b>✏️ Cosa vuoi modificare?</b></i>"
		if status == utilities.EditedChangelogs {
			text += "\n\n<i>✅ Modifiche salvate!</i>"
		}
		_, _ = c.TelegramClient.Invoke(&methods.EditMessageText{
			ChatID:    c.CacheMessage.Chat.ID,
			MessageID: c.CacheMessage.MessageID,
			Text:      text,
			ParseMode: "HTML",
			ReplyMarkup: &types.InlineKeyboardMarkup{
				InlineKeyboard: [][]types.InlineKeyboardButton{
					{
						{
							Text:         "📝 Descrizione",
							CallbackData: "release:edit_text:desc",
						},
						{
							Text:         "✍️ Note",
							CallbackData: "release:edit_text:note",
						},
					},
					{
						{
							Text:         "🔗 Link",
							CallbackData: "release:edit_text:link",
						},
						{
							Text:         "🖼 Immagine",
							CallbackData: "release:edit_text:image",
						},
					},
					{
						{
							Text:         "🔙 Indietro",
							CallbackData: "release:back_to_changelogs",
						},
					},
				},
			},
		})
		break
	case utilities.NeededImage, utilities.NeededDesc, utilities.NeededNotes, utilities.NeededLink:
		var text string
		if status == utilities.NeededImage {
			text = "<i>🖼 Invia ora la nuova immagine</i>"
		} else if status == utilities.NeededLink {
			text = "<i>🔗 Invia ora il nuovo link</i>"
		} else {
			text = "<i>💭 Invia ora il nuovo testo</i>"
		}
		_, _ = c.TelegramClient.Invoke(&methods.EditMessageText{
			ChatID:    c.CacheMessage.Chat.ID,
			MessageID: c.CacheMessage.MessageID,
			Text:      text,
			ParseMode: "HTML",
			ReplyMarkup: &types.InlineKeyboardMarkup{
				InlineKeyboard: [][]types.InlineKeyboardButton{
					{
						{
							Text:         "❌ Annulla",
							CallbackData: "release:edit_text",
						},
					},
				},
			},
		})
		break
	case utilities.ConfirmChanges:
		text := "<i>⁉️ Sei sicuro di voler applicare queste modifiche?</i>"
		_, _ = c.TelegramClient.Invoke(&methods.EditMessageText{
			ChatID:    c.CacheMessage.Chat.ID,
			MessageID: c.CacheMessage.MessageID,
			Text:      text,
			ParseMode: "HTML",
			ReplyMarkup: &types.InlineKeyboardMarkup{
				InlineKeyboard: [][]types.InlineKeyboardButton{
					{
						{
							Text:         "✅ Conferma",
							CallbackData: "release:confirm_changes",
						},
					},
					{
						{
							Text:         "❌ Annulla",
							CallbackData: "release:back_to_changelogs",
						},
					},
					{
						{
							Text:         "❌ Annulla",
							CallbackData: "release:back_to_changelogs",
						},
					},
				},
			},
		})
		break
	case utilities.ChangesConfirmed:
		text := fmt.Sprintf("<b><i>✅ OwlGram %s (%d) è stato rilasciato!</i></b>", c.CompilerClient.Bundles[0].VersionName, c.CompilerClient.Bundles[0].VersionCode)
		_, _ = c.TelegramClient.Invoke(&methods.DeleteMessage{
			ChatID:    c.CacheMessage.Chat.ID,
			MessageID: c.CacheMessage.MessageID,
		})
		_, _ = c.TelegramClient.Invoke(&methods.SendMessage{
			ChatID:    c.CacheMessage.Chat.ID,
			Text:      text,
			ParseMode: "HTML",
		})
		text += "\n\n"
		text += fmt.Sprintf("<b>• Play Console: </b><i>%s</i>", consts.GooglePlayConsoleLink)
		_, _ = c.TelegramClient.Invoke(&methods.SendMessage{
			ChatID:    consts.Tappo03UserID,
			Text:      text,
			ParseMode: "HTML",
		})
		break
	}
}
