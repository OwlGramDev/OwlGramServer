package emojipedia

import (
	"strconv"
	"strings"
)

func getHex(emoji string) string {
	emojiRaw := []rune(emoji)
	var hex []string
	for _, raw := range emojiRaw {
		hex = append(hex, strconv.FormatInt(int64(raw), 16))
	}
	return strings.Join(hex, "-")
}
