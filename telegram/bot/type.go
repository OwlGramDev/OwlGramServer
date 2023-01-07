package bot

import (
	"OwlGramServer/compiler"
	"OwlGramServer/gopy"
	"OwlGramServer/stores"
	"OwlGramServer/telegram/bot/reviews"
	"OwlGramServer/updates"
	"github.com/Squirrel-Network/gobotapi"
	"github.com/Squirrel-Network/gobotapi/types"
)

type Context struct {
	TelegramClient   *gobotapi.PollingClient
	CacheMessage     types.Message
	GoogleReviews    *reviews.Context
	CompilerClient   *compiler.Context
	StoreClient      *stores.Context
	UpdateClient     *updates.Context
	LocalizationsTmp map[string]map[string]string
	HrefTmp          string
	BannerTmp        string
	SelectedLanguage string
	ReleaseType      string
	ReleaseBase      string
	CurrentStatus    int
	SendingApks      bool
	pythonClient     *gopy.Context
}
