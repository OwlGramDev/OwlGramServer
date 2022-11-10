package updates

import (
	"OwlGramServer/consts"
	"OwlGramServer/http"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func (ctx *Context) GetChangelogUpdate(version string, lang string, ignoreChecks bool) *string {
	langStyle := ""
	switch strings.ToUpper(lang) {
	case "IT":
		langStyle = "IT"
		break
	default:
		langStyle = ""
		break
	}
	linkChannel := fmt.Sprintf("%s%s", consts.OwlGramTGChannelLink, langStyle)
	cacheId := "changelogNews"
	cacheBetaId := cacheId + "Beta"
	resultCache := ""
	updatesInfo := ctx.UpdatesDescriptor
	testVersion, _ := strconv.Atoi(version)
	if !ignoreChecks {
		if updatesInfo.Updates == nil {
			fmt.Println("CRITICAL ERROR: updatesInfo.Updates is nil")
			return nil
		}
		if updatesInfo.Updates.Beta == nil {
			fmt.Println("CRITICAL ERROR: updatesInfo.Updates.Beta is nil")
		}
		if updatesInfo.Updates.Stable == nil {
			fmt.Println("CRITICAL ERROR: updatesInfo.Updates.Stable is nil")
		}
		if testVersion > updatesInfo.Updates.Beta.VersionCode && testVersion > updatesInfo.Updates.Stable.VersionCode {
			return nil
		}
	}
	var cache *string
	if !ignoreChecks {
		cache = ctx.getCacheContent(cacheId+langStyle, version)
	}
	if cache == nil {
		result, _ := http.ExecuteRequest(linkChannel)
		tmpResult := string(result)
		cache = &tmpResult
		if !ignoreChecks {
			ctx.setCacheContent(cacheId+langStyle, version, tmpResult)
		}
	}
	if cache != nil {
		resultCache = *cache
	}
	var cacheBeta *string
	if !ignoreChecks {
		cacheBeta = ctx.getCacheContent(cacheBetaId, version)
	}
	if cacheBeta == nil {
		result, _ := http.ExecuteRequest(consts.OwlGramTGChannelBetaLink)
		tmpResult := string(result)
		cacheBeta = &tmpResult
		if !ignoreChecks {
			ctx.setCacheContent(cacheBetaId, version, tmpResult)
		}
	}
	if cacheBeta != nil {
		resultCache += "\n" + *cacheBeta
	}
	if cache != nil || cacheBeta != nil {
		r, _ := regexp.Compile(`<div class="tgme_widget_message_text js-message_text" dir="auto">(.*?)</div>`)
		for _, match := range r.FindAllSubmatch([]byte(resultCache), -1) {
			post := string(match[1])
			if strings.Contains(strings.Split(post, "<br/>")[0], version) {
				formattedString := getReadableResult(post)
				return &formattedString
			}
		}
	}
	return nil
}
