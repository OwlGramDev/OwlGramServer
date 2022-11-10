package types

type NonApprovedLanguage struct {
	LanguageCode string
	Vars         []NewVar
	LastUpdateAt int64
}
