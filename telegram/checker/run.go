package checker

import "time"

func (tg *Context) Run() {
	tg.readBackup()
	for {
		tg.doBackup()
		time.Sleep(time.Second * time.Duration(60-time.Now().Second()))
	}
}
