package types

type GoogleResult struct {
	Sentences []struct {
		Trans string `json:"trans"`
	} `json:"sentences"`
	Src string `json:"src"`
}
