package emoji

import (
	"OwlGramServer/utilities"
	"github.com/gotd/td/tg"
	"strings"
)

func (c *Context) getEmojiPacks(messages []*tg.Message) []*Pack {
	previews := make(map[string]*tg.MessageMediaDocument)
	for _, message := range messages {
		if message.Media != nil && utilities.InstanceOf(message.Media, &tg.MessageMediaDocument{}) {
			if strings.Contains(message.Message, "#preview") {
				media := message.Media.(*tg.MessageMediaDocument)
				previews[getIdentifier(media)] = media
			}
		}
	}
	packs := make(map[int]*Pack)
	for _, message := range messages {
		if message.Media != nil && utilities.InstanceOf(message.Media, &tg.MessageMediaDocument{}) {
			if strings.Contains(message.Message, "#preview") {
				continue
			}
			pack := newEmojiPack(message)
			if packs[pack.Position] == nil {
				packs[pack.Position] = pack
			} else {
				if packs[pack.Position].Version < pack.Version {
					packs[pack.Position] = pack
				} else {
					continue
				}
			}
			if preview, ok := previews[pack.Identifier]; ok {
				pack.SetPreview(preview)
			}
		}
	}
	var emojiPacks []*Pack
	for _, pack := range packs {
		emojiPacks = append(emojiPacks, pack)
	}
	return emojiPacks
}
