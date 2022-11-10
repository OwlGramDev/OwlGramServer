package updates

import (
	"OwlGramServer/updates/types"
	"strconv"
)

func (ctx *Context) getCacheContent(id string, uid string) *string {
	if ctx.cacheList == nil {
		ctx.cacheList = map[string]types.Cache{}
	}
	result, exist := ctx.cacheList[id]
	if exist {
		currUid, _ := strconv.Atoi(result.Uid)
		newUid, _ := strconv.Atoi(uid)
		if newUid > currUid {
			return nil
		}
		return &result.Content
	}
	return nil
}
