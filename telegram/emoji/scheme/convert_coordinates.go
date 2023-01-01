package scheme

import "OwlGramServer/telegram/emoji/scheme/types"

func convertCoordinates(rawCoords [][]string) map[string]*types.Coordinates {
	emojiMap := make(map[string]*types.Coordinates)
	for i, row := range rawCoords {
		for j, emoji := range row {
			emojiMap[emoji] = &types.Coordinates{X: i, Y: j}
		}
	}
	return emojiMap
}
