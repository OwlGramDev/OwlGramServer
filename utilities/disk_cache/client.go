package disk_cache

import "sync"

func Client(cacheFolder string) *Context {
	return &Context{
		data:        sync.Map{},
		cacheFolder: cacheFolder,
	}
}
