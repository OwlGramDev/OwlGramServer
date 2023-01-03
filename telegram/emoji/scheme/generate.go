package scheme

import "OwlGramServer/telegram/emoji/scheme/types"

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
	for emoji, coordsTga := range tgAScheme.Data {
		coords := tgXScheme.Data[emoji]
		if coords == nil {
			coords = tgXScheme.Data[tgAScheme.Alias[emoji]]
		}
		if coords == nil {
			coords = tgXScheme.Data[emoji]
		}
		if coords == nil {
			continue
		}
		if _, ok := newScheme.Data[coords.X]; !ok {
			newScheme.Data[coords.X] = make(map[int]*types.EmojiInfo)
		}
		newScheme.Data[coords.X][coords.Y] = &types.EmojiInfo{
			Coordinates: coordsTga,
			Emoji:       emoji,
		}
	}
	return newScheme
}
