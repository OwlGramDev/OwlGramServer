package types

type TgXScheme struct {
	Scale      float64
	SplitCount int
	Margins    [][]int
	Columns    [][]int
	Data       map[string]*Coordinates
}
