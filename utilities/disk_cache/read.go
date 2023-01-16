package disk_cache

func Read[T any](ctx *Context, key string, data *T) {
	if value, ok := ctx.data.Load(key); ok {
		*data = value.(T)
	}
}
