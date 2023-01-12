package crowdin

import (
	"OwlGramServer/consts"
	"OwlGramServer/crowdin/types"
	"OwlGramServer/http"
	"fmt"
)

func (ctx *Context) DeleteString(translationId int) {
	_ = http.ExecuteRequest(
		fmt.Sprintf("%s%s/translations/%d", consts.CrowdinApiLink, consts.CrowdinProjectId, translationId),
		http.BearerToken(consts.CrowdinAuthToken),
		http.Method("DELETE"),
		http.Retries(3),
	)
	for i, translation := range ctx.DownloadedLanguages {
		for i2, t := range translation.LanguageVars {
			if t.TranslationId == translationId {
				ctx.DownloadedLanguages[i].LanguageVars[i2] = types.TranslationInfo{
					Id:             t.Id,
					TranslationId:  t.TranslationId,
					Text:           t.ApprovedText,
					ApprovedText:   t.ApprovedText,
					IsApproved:     true,
					TranslatorName: t.TranslatorName,
					TranslatedAt:   t.TranslatedAt,
				}
				return
			}
		}
	}
}
