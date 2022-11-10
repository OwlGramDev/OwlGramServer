package utilities

import "strings"

func GetLangName(langCode string) string {
	switch strings.ToLower(langCode) {
	case "en":
		return "Inglese"
	case "it":
		return "Italiano"
	case "zh":
		return "Cinese"
	default:
		return "Inglese"
	}
}
