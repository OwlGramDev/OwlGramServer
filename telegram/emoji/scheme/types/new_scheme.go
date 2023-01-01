package types

import (
	"strconv"
	"strings"
)

type NewScheme struct {
	Scale      float64
	SplitCount int
	Margins    [][]int
	Columns    [][]int
	Data       map[int]map[int]*EmojiInfo
}

func (s *NewScheme) GetColumnDivideBy(key string) int {
	coords := strings.Split(key, "_")
	x, _ := strconv.Atoi(coords[0])
	y, _ := strconv.Atoi(coords[1])
	return s.Columns[x][y]
}
