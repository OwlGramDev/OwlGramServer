package reviews

import (
	"OwlGramServer/google_reviews"
	"github.com/Squirrel-Network/gobotapi"
)

type Context struct {
	telegramClient *gobotapi.PollingClient
	md5List        []string
	googleClient   *google_reviews.Context
}
