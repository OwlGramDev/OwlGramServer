package types

type BackupProviders struct {
	Name           string `json:"name"`
	ID             string `json:"id"`
	UnicodeVersion int    `json:"unicode_version"`
	EmojiCount     int    `json:"emoji_count"`
	MD5            string `json:"md5"`
}
