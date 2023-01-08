package emoji

import (
	"OwlGramServer/consts"
	"OwlGramServer/telegram/emoji/github"
	"OwlGramServer/telegram/emoji/scheme"
	"OwlGramServer/utilities"
	"github.com/gotd/td/tg"
	"github.com/gotd/td/tgerr"
	"sort"
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
			var applePack *Pack
			for _, p := range emojiPacks {
				if p.Identifier == "apple" {
					applePack = p
				}
			}
			sets, errBuild := github.BuildSets(applePack.Version, applePack.RawEmoji, schemeLayout, c.pythonClient)
			if errBuild == nil {
				for _, p := range sets {
					emojiPacks = append(emojiPacks, &Pack{
						DisplayName: p.DisplayName,
						Identifier:  p.Identifier,
						Preview:     p.Preview,
						Emojies:     p.Emojies,
						MD5:         p.MD5,
						Version:     p.Version,
					})
				}
			}
			sort.Slice(emojiPacks, func(i, j int) bool {
				return emojiPacks[i].DisplayName < emojiPacks[j].DisplayName
			})
			c.EmojiPacks = emojiPacks
		}
		time.Sleep(time.Minute * 30)
	}
}
