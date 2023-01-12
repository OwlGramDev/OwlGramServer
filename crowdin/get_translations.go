package crowdin

import (
	"OwlGramServer/consts"
	"OwlGramServer/crowdin/types"
	"OwlGramServer/http"
	typeRequest "OwlGramServer/http/types"
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
	var result *typeRequest.HTTPResult
	if languageId == "en" {
		result = http.ExecuteRequest(
			fmt.Sprintf("%s%s/files/4/download", consts.CrowdinApiLink, consts.CrowdinProjectId),
			http.BearerToken(consts.CrowdinAuthToken),
			http.Retries(3),
		)
	} else {
		marshal, _ := json.Marshal(types.ExportTranslations{
			TargetLanguageID:            languageId,
			Format:                      "android",
			SkipUntranslatedStrings:     false,
			SkipUntranslatedFiles:       false,
			ExportWithMinApprovalsCount: minApprovals,
		})
		result = http.ExecuteRequest(
			fmt.Sprintf("%s%s/translations/exports", consts.CrowdinApiLink, consts.CrowdinProjectId),
			http.Method("POST"),
			http.BearerToken(consts.CrowdinAuthToken),
			http.Headers(map[string]string{
				"Content-Type": "application/json",
			}),
			http.Body(marshal),
			http.Retries(3),
		)
	}
	if result.Error != nil {
		return nil, result.Error
	}
	var exportResult types.ExportResult
	err := json.Unmarshal(result.Read(), &exportResult)
	if err != nil {
		return nil, err
	}
	r := http.ExecuteRequest(
		exportResult.Data.Url,
		http.Retries(3),
	)
	compile, err := regexp.Compile("<string name=\"(.*?)\".*?>(.*?)</string>")
	if err != nil {
		return nil, err
	}
	matches := compile.FindAllStringSubmatch(r.ReadString(), -1)
	for _, match := range matches {
		translations[match[1]] = match[2]
	}
	byteSum := sha256.Sum256(r.Read())
	return &types.TranslationsInfo{
		Translations: translations,
		MD5:          hex.EncodeToString(byteSum[:]),
	}, nil
}
