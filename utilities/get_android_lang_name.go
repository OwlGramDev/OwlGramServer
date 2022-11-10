package utilities

func GetAndroidLangName(langCode string) string {
	switch langCode {
	case "ar-rSA":
		return "Arabic"
	case "be-rBY":
		return "Belarusian"
	case "ca-rES":
		return "Catalan"
	case "cs-rCZ":
		return "Czech"
	case "de-rDE":
		return "German"
	case "es-rES":
		return "Spanish"
	case "fa-rIR":
		return "Persian"
	case "fr-rFR":
		return "French"
	case "in-rID":
		return "Indonesian"
	case "it-rIT":
		return "Italian"
	case "ja-rJP":
		return "Japanese"
	case "ko-rKR":
		return "Korean"
	case "ms-rMY":
		return "Malay"
	case "nl-rNL":
		return "Dutch"
	case "pl-rPL":
		return "Polish"
	case "pt-rBR":
		return "Portuguese"
	case "ro-rRO":
		return "Romanian"
	case "ru-rRU":
		return "Russian"
	case "tr-rTR":
		return "Turkish"
	case "uk-rUA":
		return "Ukrainian"
	case "uz-rUZ":
		return "Uzbek"
	case "zh-rCN":
		return "Chinese (Simplified)"
	case "zh-rTW":
		return "Chinese (Traditional, Taiwan)"
	default:
		return "Unknown"
	}
}
