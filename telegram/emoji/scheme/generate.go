package scheme

import (
	"OwlGramServer/telegram/emoji/scheme/types"
)

func Generate() *types.NewScheme {
	tgXScheme := loadTgXScheme()
	tgAScheme := loadTgAScheme()
	newScheme := &types.NewScheme{
		Scale:      tgXScheme.Scale,
		SplitCount: tgXScheme.SplitCount,
		Margins:    tgXScheme.Margins,
		Columns:    tgXScheme.Columns,
		Data:       make(map[int]map[int]*types.EmojiInfo),
	}
	for emoji, coords := range tgXScheme.Data {
		if _, ok := newScheme.Data[coords.X]; !ok {
			newScheme.Data[coords.X] = make(map[int]*types.EmojiInfo)
		}
		if tgAScheme[emoji] == nil {
			continue
		}
		newScheme.Data[coords.X][coords.Y] = &types.EmojiInfo{
			Coordinates: tgAScheme[emoji],
			Emoji:       emoji,
		}
	}
	return newScheme
}
