package utilities

import "strings"

func GetEmojiLang(langCode string) string {
	switch strings.ToLower(langCode) {
	case "en":
		return "🇬🇧"
	case "it":
		return "🇮🇹"
	case "zh":
		return "🇨🇳"
	default:
		return "🇬🇧"
	}
}
