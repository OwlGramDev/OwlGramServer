package emoji

import (
	"OwlGramServer/emoji/emojipedia/types"
	"OwlGramServer/gopy"
	"OwlGramServer/utilities/disk_cache"
)

type Context struct {
	EmojiPacks   []*types.ProviderDescriptor
	pythonClient *gopy.Context
	cacheClient  *disk_cache.Context
}
