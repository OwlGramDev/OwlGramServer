package emoji

import (
	"OwlGramServer/telegram/emoji/scheme/types"
	"image"
)

type sprite struct {
	SectionIndex, Page int
	Rect               image.Rectangle
	Coordinates        *types.Coordinates
}
