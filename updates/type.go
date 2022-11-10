package updates

import (
	"OwlGramServer/updates/types"
)

type Context struct {
	cacheList         map[string]types.Cache
	UpdatesDescriptor *types.UpdatesDescriptor
}
