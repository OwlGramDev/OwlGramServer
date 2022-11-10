package types

type ProjectSourceStrings struct {
	Data []struct {
		Data struct {
			Id         int32  `json:"id"`
			Text       string `json:"text"`
			Identifier string `json:"identifier"`
		} `json:"data"`
	} `json:"data"`
}
