package disk_cache

func (ctx *Context) Free(key string) {
	ctx.data.Delete(key)
	ctx.synchronize()
}
