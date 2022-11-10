package types

type ExportTranslations struct {
	TargetLanguageID            string `json:"targetLanguageId"`
	Format                      string `json:"format"`
	ExportWithMinApprovalsCount int    `json:"exportWithMinApprovalsCount"`
	SkipUntranslatedStrings     bool   `json:"skipUntranslatedStrings"`
	SkipUntranslatedFiles       bool   `json:"skipUntranslatedFiles"`
}
