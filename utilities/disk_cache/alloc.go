package disk_cache

func (ctx *Context) Alloc(key string, data any) {
	ctx.data.Store(key, data)
	ctx.synchronize()
}
