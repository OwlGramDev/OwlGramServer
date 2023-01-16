package types

type EmojiMetaData struct {
	Unicode          string   `json:"unicode"`
	Cldr             string   `json:"cldr"`
	UnicodeSkinTones []string `json:"unicodeSkintones"`
}
