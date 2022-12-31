package reviews

import (
	"OwlGramServer/google_reviews"
	"github.com/Squirrel-Network/gobotapi"
)

func Client(client *gobotapi.PollingClient) *Context {
	return &Context{
		telegramClient: client,
		md5List:        []string{},
		googleClient:   google_reviews.Client(),
	}
}
