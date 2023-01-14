package emojipedia

import (
	"OwlGramServer/consts"
	"OwlGramServer/emoji/emojipedia/types"
	typesScheme "OwlGramServer/emoji/scheme/types"
	"OwlGramServer/gopy"
	"OwlGramServer/http"
	"encoding/json"
	"os"
	"path"
	"regexp"
)

func GetEmojis(scheme *typesScheme.TgAScheme, pythonClient *gopy.Context) ([]*types.ProviderDescriptor, error) {
	file, err := os.ReadFile(path.Join(consts.EmojiListConfig))
	if err != nil {
		return nil, err
	}
	var pdTemp []*types.ProviderDescriptor
	err = json.Unmarshal(file, &pdTemp)
	if err != nil {
		return nil, err
	}
	pd := make(map[string]*types.ProviderDescriptor)
	for _, v := range pdTemp {
		v.Emojis = make(map[string][]byte)
		pd[v.ID] = v
	}
	res := http.ExecuteRequest("https://emojipedia.org/beaming-face-with-smiling-eyes/")
	if res.Error != nil {
		return nil, res.Error
	}
	compile, _ := regexp.Compile(`<div.*class="vendor-image">[^<]<img.*srcset="(.*/(.*?)/\d+/).*\dx".*alt=".*?".*width="\d+".*height="\d+">`)
	linkData := compile.FindAllStringSubmatch(res.ReadString(), -1)
	for _, link := range linkData {
		if _, ok := pd[link[2]]; ok {
			pd[link[2]].Link = link[1]
		}
	}
	res = http.ExecuteRequest("https://emojipedia.org/apple/")
	if res.Error != nil {
		return nil, res.Error
	}
	compile, _ = regexp.Compile(`<img.*srcset=".*/\d+/\w+/\d+/([a-z0-9-]+_?[a-z0-9-]+)_(([^_]*)_?[^_]*)\.png\s+\dx`)
	emojiResult, err := mapEmojis(compile.FindAllStringSubmatch(res.ReadString(), -1), scheme.Data)
	if err != nil {
		return nil, err
	}
	err = downloadEmojis(pd, emojiResult)
	if err != nil {
		return nil, err
	}
	versionsCheck, err := getVersions()
	determineVersion(pd, versionsCheck)
	zipEmojis(pd, scheme.Data, pythonClient)
	for _, v := range pd {
		v.Emojis = nil
	}
	var pdTemp2 []*types.ProviderDescriptor
	for _, v := range pd {
		pdTemp2 = append(pdTemp2, v)
	}
	return pdTemp2, nil
}
