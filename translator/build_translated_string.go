package translator

import "strings"

func buildTranslatedString(original, translated string) string {
	if len(translated) > 2 {
		if strings.HasPrefix(original, "\n\n") && !strings.HasPrefix(translated, "\n\n") {
			translated = "\n\n" + translated
		} else if strings.HasPrefix(original, "\n") && !strings.HasPrefix(translated, "\n") {
			translated = "\n" + translated
		}
		if strings.HasSuffix(original, "\n\n") && !strings.HasSuffix(translated, "\n\n") {
			translated += "\n\n"
		} else if strings.HasSuffix(original, "\n") && !strings.HasSuffix(translated, "\n") {
			translated += "\n"
		}
	}
	return translated
}
