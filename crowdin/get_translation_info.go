package crowdin

import (
	"OwlGramServer/consts"
	"OwlGramServer/crowdin/types"
	"OwlGramServer/http"
	"encoding/json"
	"fmt"
)

func (ctx *Context) GetTranslationInfo(languageId string) (map[string]types.StringInfo, error) {
	r := http.ExecuteRequest(
		fmt.Sprintf("%s%s/languages/%s/translations?limit=500", consts.CrowdinApiLink, consts.CrowdinProjectId, languageId),
		http.BearerToken(consts.CrowdinAuthToken),
		http.Retries(3),
	)
	if r.Error != nil {
		return nil, r.Error
	}
	var result struct {
		Data []struct {
			Data types.StringInfo `json:"data"`
		} `json:"data"`
	}
	err := json.Unmarshal(r.Read(), &result)
	if err != nil {
		return nil, err
	}
	var resultMap = make(map[string]types.StringInfo)
	for _, v := range result.Data {
		resultMap[ctx.SourceStrings[v.Data.Id]] = v.Data
	}
	return resultMap, nil
}
