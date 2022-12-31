package emoji

import (
	"context"
	"github.com/gotd/td/tg"
)

type Context struct {
	client  *tg.Client
	context context.Context
}
