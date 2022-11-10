package updates

import "OwlGramServer/updates/types"

func (ctx *Context) setCacheContent(id string, uid string, content string) {
	if ctx.cacheList == nil {
		ctx.cacheList = map[string]types.Cache{}
	}
	ctx.cacheList[id] = types.Cache{
		Content: content,
		Uid:     uid,
	}
}
