package scheme

import (
	"OwlGramServer/consts"
	"OwlGramServer/http"
	"OwlGramServer/telegram/emoji/scheme/types"
	"encoding/json"
	"fmt"
	"regexp"
	"strings"
)

func loadTgAScheme() *types.TgAScheme {
	linkFile := fmt.Sprintf(
		"https://raw.githubusercontent.com/%s/%s/master/TMessagesProj/src/main/java/org/telegram/messenger/EmojiData.java",
		consts.GithubRepoOwner,
		consts.GithubRepo,
	)
	res, _ := http.ExecuteRequest(linkFile)
	compile, _ := regexp.Compile(`new .*\[]`)
	output := compile.ReplaceAllString(string(res), "")
	compile, _ = regexp.Compile(`public.*static.*\s(\w+).=\s(\{[^]]*}|.*);`)
	a := compile.FindAllStringSubmatch(output, -1)
	var rawMap [][]string
	var aliasOld, aliasNew []string
	tgaScheme := &types.TgAScheme{
		Alias: make(map[string]string),
	}
	for _, v := range a {
		varName := v[1]
		content := v[2]
		content = strings.Replace(content, "}", "]", -1)
		content = strings.Replace(content, "{", "[", -1)
		replaceComments, _ := regexp.Compile(`/\*.*\*/`)
		content = replaceComments.ReplaceAllString(content, "")
		replaceComments, _ = regexp.Compile(`//.*`)
		content = replaceComments.ReplaceAllString(content, "")
		if varName == "data" {
			_ = json.Unmarshal([]byte(content), &rawMap)
		} else if varName == "aliasOld" {
			_ = json.Unmarshal([]byte(content), &aliasOld)
		} else if varName == "aliasNew" {
			_ = json.Unmarshal([]byte(content), &aliasNew)
		}
	}
	tgaScheme.Data = convertCoordinates(rawMap)
	for i, v := range aliasNew {
		tgaScheme.Alias[v] = aliasOld[i]
	}
	return tgaScheme
}
