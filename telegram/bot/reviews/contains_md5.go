package reviews

func (ctx *Context) containsMD5(md5 string) bool {
	for _, a := range ctx.md5List {
		if a == md5 {
			return true
		}
	}
	return false
}
