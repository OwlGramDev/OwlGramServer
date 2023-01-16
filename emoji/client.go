package emoji

import (
	"OwlGramServer/gopy"
	"OwlGramServer/utilities/disk_cache"
)

func Client(pythonClient *gopy.Context, cacheClient *disk_cache.Context) *Context {
	return &Context{
		pythonClient: pythonClient,
		cacheClient:  cacheClient,
	}
}
