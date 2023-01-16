package emoji

import (
	"OwlGramServer/consts"
	"OwlGramServer/emoji/types"
	"encoding/json"
	"os"
	"path"
)

func (c *Context) backup() {
	_ = os.MkdirAll(consts.CacheFolderEmojis, os.ModePerm)
	for _, pack := range c.EmojiPacks {
		filePath := path.Join(consts.CacheFolderEmojis, LegacyID(pack.ID, pack.GetID()))
		_ = os.WriteFile(filePath+".png", pack.Preview, os.ModePerm)
		_ = os.WriteFile(filePath+".zip", pack.EmojiZip, os.ModePerm)
	}
	var backupProviders []types.BackupProviders
	for _, provider := range c.EmojiPacks {
		backupProviders = append(backupProviders, types.BackupProviders{
			Name:           provider.Name,
			ID:             provider.ID,
			UnicodeVersion: provider.UnicodeVersion,
			EmojiCount:     provider.EmojiCount,
			MD5:            provider.MD5,
		})
	}
	res, _ := json.Marshal(backupProviders)
	_ = os.WriteFile(path.Join(consts.CacheFolderEmojis, "providers.json"), res, os.ModePerm)
}
