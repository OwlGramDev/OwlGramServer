package types

import (
	typesScheme "OwlGramServer/emoji/scheme/types"
)

type EmojiRaw struct {
	Coordinates *typesScheme.Coordinates
	EmojiHex    string
	EmojiName   string
}
