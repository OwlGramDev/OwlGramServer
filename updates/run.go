package updates

import (
	"time"
)

func (ctx *Context) Run() {
	for {
		ctx.readUpdateInfo()
		time.Sleep(time.Second * 5)
	}
}
