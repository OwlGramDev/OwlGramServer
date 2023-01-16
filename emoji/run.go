package emoji

import (
	"OwlGramServer/emoji/emojipedia"
	"OwlGramServer/emoji/scheme"
	"OwlGramServer/utilities/disk_cache"
	"sort"
	"time"
)

func (c *Context) Run() {
	for {
		disk_cache.Read(c.cacheClient, "emoji_packs", &c.EmojiPacks)
		res, _ := emojipedia.GetEmojis(scheme.LoadTgAScheme(), c.pythonClient)
		sort.Slice(res, func(i, j int) bool {
			return res[i].Name < res[j].Name
		})
		c.EmojiPacks = res
		c.cacheClient.Alloc("emoji_packs", c.EmojiPacks)
		time.Sleep(time.Hour * 5)
	}
}
