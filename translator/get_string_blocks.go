package translator

import (
	"math"
	"strings"
)

func getStringBlocks(query string, maxBlockSize int) []string {
	var blocks []string
	for len(query) > maxBlockSize {
		maxBlockStr := query[:maxBlockSize]
		n := strings.LastIndex(maxBlockStr, "\n\n")
		if n == -1 {
			n = strings.LastIndex(maxBlockStr, "\n")
		}
		if n == -1 {
			n = strings.LastIndex(maxBlockStr, ". ")
		}
		if n == -1 {
			n = int(math.Min(float64(len(query)), float64(maxBlockSize)))
		}
		blocks = append(blocks, query[:n])
		query = query[n:]
	}
	if len(query) > 0 {
		blocks = append(blocks, query)
	}
	return blocks
}
