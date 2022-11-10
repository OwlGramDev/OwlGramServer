package crowdin

import (
	"OwlGramServer/consts"
	"OwlGramServer/crowdin/types"
	"OwlGramServer/http"
	"encoding/json"
	"fmt"
)

func (ctx *Context) ApproveString(translationId int) {
	marshal, _ := json.Marshal(map[string]int{
		"translationId": translationId,
	})
	_, err := http.ExecuteRequest(
		fmt.Sprintf("%s%s/approvals", consts.CrowdinApiLink, consts.CrowdinProjectId),
		http.Method("POST"),
		http.Body(marshal),
		http.BearerToken(consts.CrowdinAuthToken),
		http.Headers(map[string]string{
			"Content-Type": "application/json",
		}),
		http.Retries(3),
	)
	if err != nil {
		return
	}
	for i, translation := range ctx.DownloadedLanguages {
		for i2, t := range translation.LanguageVars {
			if t.TranslationId == translationId {
				ctx.DownloadedLanguages[i].LanguageVars[i2] = types.TranslationInfo{
					Id:             t.Id,
					TranslationId:  t.TranslationId,
					Text:           t.Text,
					ApprovedText:   t.Text,
					IsApproved:     true,
					TranslatorName: t.TranslatorName,
					TranslatedAt:   t.TranslatedAt,
				}
				return
			}
		}
	}
}
