package utilities

import (
	"strconv"
	"strings"
)

func GetHexFromEmoji(emoji, split string, withVariant bool) string {
	emojiRaw := []rune(emoji)
	var hex []string
	for _, raw := range emojiRaw {
		tmpHex := strconv.FormatInt(int64(raw), 16)
		if !withVariant && tmpHex == "fe0f" {
			continue
		}
		hex = append(hex, tmpHex)
	}
	return strings.Join(hex, split)
}
