package updates

import "strings"

func GetLocalization(lang string, localizations map[string]map[string]string) map[string]string {
	lang = strings.ToLower(lang)
	if localizations[lang] != nil {
		return localizations[lang]
	} else {
		return localizations["en"]
	}
}
