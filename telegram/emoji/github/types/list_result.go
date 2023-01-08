package types

import "time"

type FileDescriptor struct {
	Name         string
	ID           string
	DownloadURL  string
	LastUpdate   time.Time
	Content      []byte
	EmojiSprites map[int]map[int][]byte
}
