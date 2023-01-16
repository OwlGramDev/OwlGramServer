package utilities

import (
	"strconv"
	"strings"
)

func GetEmojiFromHex(hex, join string, withVariant bool) string {
	hexArr := strings.Split(hex, join)
	var emojiArr []rune
	for _, h := range hexArr {
		if !withVariant && h == "fe0f" {
			continue
		}
		i, _ := strconv.ParseInt(h, 16, 32)
		emojiArr = append(emojiArr, rune(i))
	}
	return string(emojiArr)
}
