package emojipedia

import (
	"OwlGramServer/emoji/emojipedia/types"
	typesScheme "OwlGramServer/emoji/scheme/types"
	"fmt"
	"strings"
)

func mapEmojis(emojis [][]string, mapScheme map[string]*typesScheme.Coordinates) (map[string]*types.EmojiRaw, error) {
	emojiResult := make(map[string]*types.EmojiRaw)
	for emoji, dataEmoji := range mapScheme {
		var emojiFound bool
		for _, emojiData := range emojis {
			emojiRemote := strings.ReplaceAll(emojiData[3], "-fe0f", "")
			emojiLocal := strings.ReplaceAll(getHex(emoji), "-fe0f", "")
			emojiName := emojiData[1]
			if emojiLocal == emojiRemote {
				if _, ok := emojiResult[emoji]; ok {
					return nil, fmt.Errorf("emoji %s already exist", emojiRemote)
				}
				emojiFound = true
				emojiResult[emoji] = &types.EmojiRaw{
					Coordinates: dataEmoji,
					EmojiHex:    emojiData[2],
					EmojiName:   emojiName,
				}
				break
			}
		}
		if !emojiFound {
			return nil, fmt.Errorf("emoji %s not found", getHex(emoji))
		}
	}
	return emojiResult, nil
}
