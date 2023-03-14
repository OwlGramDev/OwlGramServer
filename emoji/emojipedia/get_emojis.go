package emojipedia

import (
	"OwlGramServer/consts"
	"OwlGramServer/emoji/emojipedia/types"
	"OwlGramServer/emoji/github"
	typesScheme "OwlGramServer/emoji/scheme/types"
	"OwlGramServer/gopy"
	"OwlGramServer/http"
	httpTypes "OwlGramServer/http/types"
	"encoding/json"
	"fmt"
	"os"
	"path"
	"regexp"
	"strings"
)

func GetEmojis(scheme *typesScheme.TgAScheme, pythonClient *gopy.Context) ([]*types.ProviderDescriptor, error) {
	file, err := os.ReadFile(path.Join(consts.EmojiListConfig))
	if err != nil {
		return nil, err
	}
	var pd []*types.ProviderDescriptor
	err = json.Unmarshal(file, &pd)
	if err != nil {
		return nil, err
	}
	for _, v := range pd {
		v.Emojis = make(map[string][]byte)
	}
	res := http.ExecuteRequest("https://emojipedia.org/beaming-face-with-smiling-eyes/")
	if res.Error != nil {
		return nil, res.Error
	}
	compile, _ := regexp.Compile(`<div.*class="vendor-image">[^<]<img.*srcset="(.*/(.*?)/\d+/).*\dx".*alt=".*?".*width="\d+".*height="\d+">`)
	linkData := compile.FindAllStringSubmatch(res.ReadString(), -1)
	for _, v := range pd {
		for _, link := range linkData {
			if v.GetID() == link[2] {
				if v.HaveVariant() {
					tmpLink := strings.Index(link[1], v.GetID())
					tmpHtml := http.ExecuteRequest(fmt.Sprintf("https://emojipedia.org/%s/beaming-face-with-smiling-eyes/", v.ID))
					if tmpHtml.Error != nil {
						return nil, tmpHtml.Error
					}
					tmpRegex, _ := regexp.Compile(`<img src=".*?/\d+/.*?/(\d+)/.*"\s*srcset=".*2x".*?>`)
					tmpData := tmpRegex.FindAllStringSubmatch(tmpHtml.ReadString(), -1)
					if tmpLink != -1 && len(tmpData) > 0 {
						v.Link = fmt.Sprintf("%s%s/%s/", link[1][:tmpLink], v.GetID(), tmpData[0][1])
					} else {
						return nil, fmt.Errorf("can't find link for %s", v.ID)
					}
				} else {
					v.Link = link[1]
				}
				v.Link = strings.ReplaceAll(v.Link, "thumbs/240", "thumbs/120")
				break
			}
		}
	}
	appleRes := http.ExecuteRequest("https://emojipedia.org/apple/")
	for _, v := range pd {
		var res *httpTypes.HTTPResult
		if v.HaveVariant() {
			res = http.ExecuteRequest(fmt.Sprintf("https://emojipedia.org/%s/", v.ID))
		} else {
			res = appleRes
		}
		if res.Error != nil {
			return nil, res.Error
		}
		compile, _ = regexp.Compile(`<img.*srcset=".*/\d+/\w+/\d+/([a-z0-9-]+_?[a-z0-9-]+)_(([^_]*)_?[^_]*)\.png\s+\dx`)
		emojiResult, err := mapEmojis(compile.FindAllStringSubmatch(res.ReadString(), -1), scheme.Data)
		if err != nil {
			return nil, err
		}
		v.EmojiMap = emojiResult
	}

	err = downloadEmojis(pd)
	if err != nil {
		return nil, err
	}
	for _, v := range pd {
		if v.GetID() == "apple" {
			v.Emojis, err = github.DownloadApple(scheme)
			if err != nil {
				return nil, err
			}
			break
		}
	}

	// Enhance Fluent Emojis with GitHub Emojis
	zip, err := github.DownloadZip(scheme)
	if err != nil {
		return nil, err
	}
	for _, v := range pd {
		if v.GetID() == "microsoft-teams" {
			for k, c := range zip {
				v.Emojis[k] = c
			}
			break
		}
	}
	versionsCheck, err := getVersions()
	determineVersion(pd, versionsCheck)
	zipEmojis(pd, scheme.Data, pythonClient)
	for _, v := range pd {
		v.Emojis = nil
	}
	return pd, nil
}
