package disk_cache

import "sync"

type Context struct {
	data        sync.Map
	cacheFolder string
	mutex       sync.RWMutex
}
