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
	for _, coordsTgx := range tgXScheme.Data {
		if _, ok := newScheme.Data[coordsTgx.X]; !ok {
			newScheme.Data[coordsTgx.X] = make(map[int]*types.EmojiInfo)
		}
		if _, ok := newScheme.Data[coordsTgx.X][coordsTgx.Y]; !ok {
			newScheme.Data[coordsTgx.X][coordsTgx.Y] = nil
		}
	}
	return newScheme
}
