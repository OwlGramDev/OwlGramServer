package types

type PacksInfo struct {
	Name       string `json:"name"`
	Version    int    `json:"version"`
	ID         string `json:"id"`
	Preview    string `json:"preview"`
	File       string `json:"file"`
	FileSize   int    `json:"file_size"`
	EmojiCount int    `json:"emoji_count"`
	MD5        string `json:"md5"`
}
