package github

import (
	"OwlGramServer/gopy"
	"OwlGramServer/telegram/emoji/github/types"
	typeScheme "OwlGramServer/telegram/emoji/scheme/types"
)

func BuildSets(appleEmoji map[int]map[int][]byte, schemeLayout *typeScheme.NewScheme, pythonClient *gopy.Context) ([]*types.PackTMP, error) {
	data, err := getSprites(schemeLayout, pythonClient)
	if err != nil {
		return nil, err
	}
	return buildSprites(appleEmoji, schemeLayout, data)
}
