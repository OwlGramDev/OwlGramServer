package bot

import (
	"OwlGramServer/consts"
	"github.com/Squirrel-Network/gobotapi/filters"
)

func (c *Context) Run() {
	c.TelegramClient.OnMessage(filters.Filter(c.handleMessage, filters.UserID(consts.AuthorizedUsers...)))
	c.TelegramClient.OnCallbackQuery(filters.Filter(c.handleCallbackQuery, filters.UserID(consts.AuthorizedUsers...)))
	c.handleCommands()
	if !consts.IsDebug {
		go c.GoogleReviews.Run()
	}
	c.TelegramClient.Start()
}
