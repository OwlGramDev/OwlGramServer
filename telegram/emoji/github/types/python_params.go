package types

import "OwlGramServer/telegram/emoji/scheme/types"

type PythonParams struct {
	EmojiFile []byte                           `json:"emoji_file"`
	Glyphs    map[int]map[int]*types.EmojiInfo `json:"glyphs"`
}
