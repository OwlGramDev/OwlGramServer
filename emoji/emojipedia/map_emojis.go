package emojipedia

import (
	"OwlGramServer/emoji/emojipedia/types"
	typesScheme "OwlGramServer/emoji/scheme/types"
	"OwlGramServer/utilities"
	"fmt"
	"strings"
)

func mapEmojis(emojis [][]string, mapScheme map[string]*typesScheme.Coordinates) (map[string]*types.EmojiRaw, error) {
	emojiResult := make(map[string]*types.EmojiRaw)
	for emoji, dataEmoji := range mapScheme {
		for _, emojiData := range emojis {
			emojiRemote := strings.ReplaceAll(emojiData[3], "-fe0f", "")
			emojiLocal := utilities.GetHexFromEmoji(emoji, "-", false)
			emojiName := emojiData[1]
			if emojiLocal == emojiRemote {
				if _, ok := emojiResult[emoji]; ok {
					return nil, fmt.Errorf("emoji %s already exist", emojiRemote)
				}
				emojiResult[emoji] = &types.EmojiRaw{
					Coordinates: dataEmoji,
					EmojiHex:    emojiData[2],
					EmojiName:   emojiName,
				}
				break
			}
		}
	}
	return emojiResult, nil
}
