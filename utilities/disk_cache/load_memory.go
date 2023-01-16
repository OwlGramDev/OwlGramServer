package disk_cache

import (
	"OwlGramServer/utilities/bytecode"
	"os"
)

func (ctx *Context) LoadMemory() {
	if _, err := os.Stat(ctx.cacheFolder); !os.IsNotExist(err) {
		file, err := os.ReadFile(ctx.cacheFolder)
		if err != nil {
			return
		}
		var cacheTmp map[string]any
		err = bytecode.Unmarshal(file, &cacheTmp)
		if err != nil {
			return
		}
		for key, value := range cacheTmp {
			ctx.data.Store(key, value)
		}
	}
}
