package emoji

import (
	"OwlGramServer/consts"
	emojiPedia "OwlGramServer/emoji/emojipedia/types"
	"OwlGramServer/emoji/types"
	"encoding/json"
	"os"
	"path"
)

func (c *Context) restore() {
	var backupProviders []types.BackupProviders
	file, _ := os.ReadFile(path.Join(consts.CacheFolderEmojis, "providers.json"))
	_ = json.Unmarshal(file, &backupProviders)
	for _, provider := range backupProviders {
		pack := &emojiPedia.ProviderDescriptor{
			Name:           provider.Name,
			ID:             provider.ID,
			UnicodeVersion: provider.UnicodeVersion,
			EmojiCount:     provider.EmojiCount,
			MD5:            provider.MD5,
		}
		filePath := path.Join(consts.CacheFolderEmojis, LegacyID(pack.ID, pack.GetID()))
		preview, _ := os.ReadFile(filePath + ".png")
		emojiZip, _ := os.ReadFile(filePath + ".zip")
		pack.Preview = preview
		pack.EmojiZip = emojiZip
		c.EmojiPacks = append(c.EmojiPacks, pack)
	}
}
