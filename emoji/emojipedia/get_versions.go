package emojipedia

import (
	"OwlGramServer/http"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func getVersions() (map[int][]string, error) {
	res := http.ExecuteRequest("https://emojipedia.org/emoji-12.0/")
	if res.Error != nil {
		return nil, res.Error
	}
	compile := regexp.MustCompile(`<li>[^<]<a href="/emoji-([0-9.]+)/"`)
	versions := compile.FindAllStringSubmatch(res.ReadString(), -1)
	versionsMap := make(map[int][]string)
	for _, v := range versions {
		version := v[1]
		versionInt, _ := strconv.Atoi(strings.Split(version, ".")[0])
		r := http.ExecuteRequest(fmt.Sprintf("https://emojipedia.org/emoji-%s/", version))
		if r.Error != nil {
			return nil, r.Error
		}
		htmlRes := r.ReadString()
		startCut := strings.Index(htmlRes, fmt.Sprintf(`<h2>New Emojis in Version %s</h2>`, version))
		endCut := strings.Index(htmlRes, `<h2>Related</h2>`)
		htmlRes = htmlRes[startCut:endCut]
		emojiAvailable := regexp.MustCompile(`<li><a\s*href="/[a-z0-9-]+/"><span\s*class="emoji">(.*?)</span>.*</a></li>`)
		emojis := emojiAvailable.FindAllStringSubmatch(htmlRes, -1)
		for _, e := range emojis {
			versionsMap[versionInt] = append(versionsMap[versionInt], e[1])
		}
	}
	return versionsMap, nil
}
