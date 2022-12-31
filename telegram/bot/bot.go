package bot

import (
	"OwlGramServer/compiler"
	"OwlGramServer/consts"
	"OwlGramServer/stores"
	"OwlGramServer/telegram/bot/reviews"
	"OwlGramServer/updates"
	"github.com/Squirrel-Network/gobotapi"
	"github.com/Squirrel-Network/gobotapi/logger"
	"os"
)

func Bot(updateClient *updates.Context) *Context {
	client := gobotapi.NewClient(consts.BotToken)
	client.SleepThreshold = 120
	if consts.IsDebug {
		client.LoggingLevel = logger.Error
	} else {
		client.LoggingLevel = logger.Silent
	}
	if _, err := os.Stat(consts.BotImagesCache); os.IsNotExist(err) {
		_ = os.Mkdir(consts.BotImagesCache, 0775)
	} else {
		_ = os.RemoveAll(consts.BotImagesCache)
		_ = os.Mkdir(consts.BotImagesCache, 0775)
	}
	return &Context{
		TelegramClient: client,
		GoogleReviews:  reviews.Client(client),
		CompilerClient: compiler.Client(),
		StoreClient:    stores.Client(updateClient),
		UpdateClient:   updateClient,
	}
}
