package reviews

import (
	"OwlGramServer/consts"
	"OwlGramServer/translator"
	"OwlGramServer/utilities"
	"fmt"
	"github.com/Squirrel-Network/gobotapi/methods"
	"golang.org/x/text/language"
	"golang.org/x/text/language/display"
	"sort"
	"strconv"
	"time"
)

func (ctx *Context) Run() {
	ctx.readBackup()
	for {
		reviews := ctx.googleClient.GetAllReview()
		if reviews != nil {
			reviewsList := *reviews
			sort.SliceStable(reviewsList, func(i, j int) bool {
				return reviewsList[i].LastEdit < reviewsList[j].LastEdit
			})
			for reviewIndex := range reviewsList {
				review := reviewsList[reviewIndex]
				currMD5 := review.GetMD5()
				if !ctx.containsMD5(currMD5) {
					notificationMessage := "<b>⭐️ Nuova Recensione</b>\n"
					if len(review.AuthorName) > 0 {
						notificationMessage += "<b>• Utente:</b> " + review.AuthorName + "\n"
					}
					notificationMessage += "<b>• Stelle:</b> <code>" + strconv.Itoa(int(review.StarRating)) + "</code>\n"
					if len(review.Text) > 0 {
						notificationMessage += "<b>• Testo:</b>" + review.Text + "\n"
						res := translator.Translate(review.Text, "it")
						if res != nil && res.SourceLanguage != "it" {
							italian := display.Italian.Languages()
							source := language.MustParse(res.SourceLanguage)
							to := language.MustParse("it")
							notificationMessage += fmt.Sprintf("\n<b>Traduzione %s > %s </b>\n", utilities.Capitalize(italian.Name(source)), utilities.Capitalize(italian.Name(to)))
							notificationMessage += fmt.Sprintf("<i>%s</i>\n", res.Translation)
						}
					}
					notificationMessage += "\n<b>🤖 Info sul dispositivo</b>\n"
					if review.DeviceModel != nil {
						notificationMessage += "<b>• Modello:</b> " + *review.DeviceModel + "\n"
						notificationMessage += "<b>• CPU:</b> " + *review.CPUVendor + "\n"
					}
					var androidName string
					if review.AndroidSDK >= 0 {
						if int(review.AndroidSDK)-1 < len(consts.AndroidVersionMap) {
							androidName = consts.AndroidVersionMap[int(review.AndroidSDK)-1]
						} else {
							androidName = fmt.Sprintf("??? (%d)", review.AndroidSDK)
						}
					} else {
						androidName = "Sconosciuto"
					}
					notificationMessage += "<b>• OS Android:</b> <code>" + androidName + "</code>\n"
					if review.AppVersionCode != 0 {
						abi := "unknown"
						switch review.AppVersionCode % 10 {
						case 1:
						case 3:
							abi = "arm-v7a"
							break
						case 2:
						case 4:
							abi = "x86"
							break
						case 5:
						case 7:
							abi = "arm64-v8a"
							break
						case 6:
						case 8:
							abi = "x86_64"
							break
						case 0:
						case 9:
							abi = "universal"
							break
						}
						notificationMessage += "<b>• Architettura:</b> <code>" + abi + "</code>\n"
						notificationMessage += "<b>• Versione App:</b> <code>" + review.AppVersionName + " (" + strconv.Itoa(int(review.AppVersionCode)/10) + ")</code>\n"
					}
					notificationMessage += "\n#recensione"
					_, _ = ctx.telegramClient.Invoke(&methods.SendMessage{
						ChatID:    consts.StaffGroupID,
						Text:      notificationMessage,
						ParseMode: "HTML",
					})
					ctx.md5List = append(ctx.md5List, currMD5)
				}
			}
		}
		ctx.doBackup()
		time.Sleep(time.Second * time.Duration(30-time.Now().Second()))
	}
}
