package disk_cache

import (
	"OwlGramServer/utilities/bytecode"
	"os"
)

func (ctx *Context) synchronize() {
	ctx.mutex.Lock()
	go func() {
		cacheTmp := make(map[string]any)
		ctx.data.Range(func(key, value interface{}) bool {
			cacheTmp[key.(string)] = value
			return true
		})
		marshal, err := bytecode.Marshal(cacheTmp)
		if err != nil {
			return
		}
		_ = os.WriteFile(ctx.cacheFolder, marshal, os.ModePerm)
		ctx.mutex.Unlock()
	}()
}
