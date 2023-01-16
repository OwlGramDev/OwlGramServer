package emoji

import (
	"OwlGramServer/emoji/emojipedia/types"
	"OwlGramServer/gopy"
)

type Context struct {
	EmojiPacks   []*types.ProviderDescriptor
	pythonClient *gopy.Context
}
