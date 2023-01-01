package emoji

import (
	"OwlGramServer/consts"
	"OwlGramServer/telegram/emoji/scheme"
	"OwlGramServer/utilities"
	"github.com/gotd/td/tg"
	"github.com/gotd/td/tgerr"
	"sync"
	"time"
)

func (c *Context) Run() {
	result, err := c.client.ContactsResolveUsername(c.context, consts.TelegramXChannel)
	if err != nil {
		panic(err)
	}
	channel := result.Chats[0].(*tg.Channel)
	for {
		history, err := c.client.MessagesGetHistory(c.context, &tg.MessagesGetHistoryRequest{
			Peer: &tg.InputPeerChannel{
				ChannelID:  channel.ID,
				AccessHash: channel.AccessHash,
			},
			Limit: 100,
		})
		schemeLayout := scheme.Generate()
		floodTest, isFloodWait := tgerr.AsFloodWait(err)
		if isFloodWait {
			time.Sleep(time.Duration(floodTest.Nanoseconds()))
			continue
		} else if err == nil {
			var messages []*tg.Message
			for _, rawMess := range history.(*tg.MessagesChannelMessages).Messages {
				if utilities.InstanceOf(rawMess, &tg.Message{}) {
					messages = append(messages, rawMess.(*tg.Message))
				}
			}
			emojiPacks := c.getEmojiPacks(messages)
			var waitUntil sync.WaitGroup
			for _, p := range emojiPacks {
				waitUntil.Add(1)
				go func(p *Pack) {
					p.Download(c)
					p.BuildPack(schemeLayout)
					waitUntil.Done()
				}(p)
			}
			waitUntil.Wait()
			c.EmojiPacks = emojiPacks
		}
		time.Sleep(time.Minute * 30)
	}
}
