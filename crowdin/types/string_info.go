package types

type StringInfo struct {
	Id            int    `json:"stringId"`
	TranslationId int    `json:"translationId"`
	CreatedAt     string `json:"createdAt"`
	User          struct {
		FullName string `json:"fullName"`
	} `json:"user"`
}
