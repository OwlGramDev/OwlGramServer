package types

type TgAScheme struct {
	Data  map[string]*Coordinates
	Alias map[string]string
}

func (t *TgAScheme) GetCoordinates(emoji string) *Coordinates {
	if t.Alias[emoji] != "" {
		emoji = t.Alias[emoji]
	}
	return t.Data[emoji]
}

func (t *TgAScheme) GetEmoji(coordinates Coordinates) string {
	for emoji, coord := range t.Data {
		if coord.X == coordinates.X && coord.Y == coordinates.Y {
			return emoji
		}
	}
	return ""
}
