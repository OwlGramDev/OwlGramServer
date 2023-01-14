package emoji

import (
	"OwlGramServer/emoji/emojipedia"
	"OwlGramServer/emoji/scheme"
	"sort"
	"time"
)

func (c *Context) Run() {
	for {
		res, _ := emojipedia.GetEmojis(scheme.LoadTgAScheme(), c.pythonClient)
		sort.Slice(res, func(i, j int) bool {
			return res[i].Name < res[j].Name
		})
		c.EmojiPacks = res
		time.Sleep(time.Hour * 5)
	}
}
