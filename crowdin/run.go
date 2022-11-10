package crowdin

import "time"

func (ctx *Context) Run() {
	for {
		ctx.downloadLanguages()
		t := time.Now()
		time.Sleep(time.Second * time.Duration(60-t.Second()))
	}
}
