package types

type TranslationInfo struct {
	Id             int
	TranslationId  int
	Text           string
	ApprovedText   string
	IsApproved     bool
	TranslatorName string
	TranslatedAt   string
}
