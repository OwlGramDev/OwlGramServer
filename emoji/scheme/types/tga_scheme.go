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
