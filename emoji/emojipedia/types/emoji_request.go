package types

import "OwlGramServer/http/types"

type EmojiRequest struct {
	Provider   string
	EmojiLink  string
	Emoji      string
	HttpResult *types.HTTPResult
	Content    []byte
}
