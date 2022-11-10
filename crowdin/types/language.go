package types

type Language struct {
	MD5          string                     `json:"md5"`
	LanguageCode string                     `json:"lang"`
	LanguageVars map[string]TranslationInfo `json:"lang_vars"`
}
