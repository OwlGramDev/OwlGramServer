package emoji

import (
	"OwlGramServer/consts"
	"OwlGramServer/emoji/emojipedia"
	"OwlGramServer/emoji/scheme"
	"sort"
	"time"
)

func (c *Context) Run() {
	var isRunning bool
	for {
		c.restore()
		if len(c.EmojiPacks) == 0 || isRunning || consts.IsDebug {
			res, _ := emojipedia.GetEmojis(scheme.LoadTgAScheme(), c.pythonClient)
			sort.Slice(res, func(i, j int) bool {
				return res[i].Name < res[j].Name
			})
			c.EmojiPacks = res
		}
		isRunning = true
		c.backup()
		time.Sleep(time.Hour * 5)
	}
}
