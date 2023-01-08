package emoji

import (
	"OwlGramServer/gopy"
	"context"
	"github.com/gotd/td/tg"
)

type Context struct {
	client       *tg.Client
	context      context.Context
	pythonClient *gopy.Context
	EmojiPacks   []*Pack
}
