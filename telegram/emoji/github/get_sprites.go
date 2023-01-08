package github

import (
	"OwlGramServer/consts"
	"OwlGramServer/gopy"
	"OwlGramServer/telegram/emoji/github/types"
	typeScheme "OwlGramServer/telegram/emoji/scheme/types"
	"encoding/json"
	"path"
	"sync"
)

func getSprites(schemeLayout *typeScheme.NewScheme, pythonClient *gopy.Context) ([]*types.FileDescriptor, error) {
	emojiSets := GetLatestEmojis()
	var waiGroup sync.WaitGroup
	var foundError error
	for _, emojiSet := range emojiSets {
		waiGroup.Add(1)
		go func(emojiSet *types.FileDescriptor) {
			defer waiGroup.Done()
			res, err := pythonClient.Run(path.Join(consts.PythonLibPillow, "__main__.py"), types.PythonParams{
				EmojiFile: emojiSet.Content,
				Glyphs:    schemeLayout.Data,
			})
			if err != nil {
				foundError = err
				return
			}
			err = json.Unmarshal(res, &emojiSet.EmojiSprites)
			if err != nil {
				foundError = err
				return
			}
		}(emojiSet)
	}
	waiGroup.Wait()
	if foundError != nil {
		return nil, foundError
	}
	return emojiSets, nil
}
