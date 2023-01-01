package scheme

import (
	"OwlGramServer/consts"
	"OwlGramServer/http"
	"OwlGramServer/telegram/emoji/scheme/types"
	"encoding/json"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func loadTgXScheme() *types.TgXScheme {
	linkFile := fmt.Sprintf(
		"https://raw.githubusercontent.com/%s/%s/main/app/src/main/java/org/thunderdog/challegram/tool/EmojiCode.java",
		consts.GithubRepoOwnerTgX,
		consts.GithubRepoTgX,
	)
	res, _ := http.ExecuteRequest(linkFile)
	compile, _ := regexp.Compile(`new .*\[]`)
	output := compile.ReplaceAllString(string(res), "")
	compile, _ = regexp.Compile(`public.*static.*\s(\w+).=\s(\{[^]]*}|.*);`)
	a := compile.FindAllStringSubmatch(output, -1)
	tgxScheme := &types.TgXScheme{}
	var rawMap [][]string
	for _, v := range a {
		varName := v[1]
		content := v[2]
		content = strings.Replace(content, "}", "]", -1)
		content = strings.Replace(content, "{", "[", -1)
		replaceComments, _ := regexp.Compile(`/\*.*\*/`)
		content = replaceComments.ReplaceAllString(content, "")
		replaceComments, _ = regexp.Compile(`//.*`)
		content = replaceComments.ReplaceAllString(content, "")
		if varName == "COLUMNS" {
			_ = json.Unmarshal([]byte(content), &tgxScheme.Columns)
		} else if varName == "DATA" {
			_ = json.Unmarshal([]byte(content), &rawMap)
		} else if varName == "MARGINS" {
			_ = json.Unmarshal([]byte(content), &tgxScheme.Margins)
		} else if varName == "SPLIT_COUNT" {
			tgxScheme.SplitCount, _ = strconv.Atoi(content)
		} else if varName == "SCALE" {
			content = strings.Replace(content, "f", "", -1)
			tgxScheme.Scale, _ = strconv.ParseFloat(content, 64)
		}
	}
	tgxScheme.Data = convertCoordinates(rawMap)
	return tgxScheme
}
