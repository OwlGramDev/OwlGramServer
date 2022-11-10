package types

type ProjectInfo struct {
	Data struct {
		TargetLanguages []LanguageData `json:"targetLanguages"`
	} `json:"data"`
}
