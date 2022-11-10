package crowdin

import (
	"OwlGramServer/consts"
	"OwlGramServer/crowdin/types"
	"OwlGramServer/http"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"regexp"
)

func (ctx *Context) GetTranslationsExported(languageId string, onlyApproved bool) (*types.TranslationsInfo, error) {
	translations := make(map[string]string)
	minApprovals := 0
	if onlyApproved {
		minApprovals = 1
	}
	var request []byte
	if languageId == "en" {
		r, err := http.ExecuteRequest(
			fmt.Sprintf("%s%s/files/4/download", consts.CrowdinApiLink, consts.CrowdinProjectId),
			http.BearerToken(consts.CrowdinAuthToken),
			http.Retries(3),
		)
		request = r
		if err != nil {
			return nil, err
		}
	} else {
		marshal, err := json.Marshal(types.ExportTranslations{
			TargetLanguageID:            languageId,
			Format:                      "android",
			SkipUntranslatedStrings:     false,
			SkipUntranslatedFiles:       false,
			ExportWithMinApprovalsCount: minApprovals,
		})
		if err != nil {
			return nil, err
		}
		request, err = http.ExecuteRequest(
			fmt.Sprintf("%s%s/translations/exports", consts.CrowdinApiLink, consts.CrowdinProjectId),
			http.Method("POST"),
			http.BearerToken(consts.CrowdinAuthToken),
			http.Headers(map[string]string{
				"Content-Type": "application/json",
			}),
			http.Body(marshal),
			http.Retries(3),
		)
		if err != nil {
			return nil, err
		}
	}
	var exportResult types.ExportResult
	err := json.Unmarshal(request, &exportResult)
	if err != nil {
		return nil, err
	}
	request, _ = http.ExecuteRequest(
		exportResult.Data.Url,
		http.Retries(3),
	)
	compile, err := regexp.Compile("<string name=\"(.*?)\".*?>(.*?)</string>")
	if err != nil {
		return nil, err
	}
	matches := compile.FindAllStringSubmatch(string(request), -1)
	for _, match := range matches {
		translations[match[1]] = match[2]
	}
	byteSum := sha256.Sum256(request)
	return &types.TranslationsInfo{
		Translations: translations,
		MD5:          hex.EncodeToString(byteSum[:]),
	}, nil
}
