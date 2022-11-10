package crowdin

import (
	"OwlGramServer/crowdin/types"
	"sort"
	"time"
)

func (ctx *Context) GetAllApprovedStrings() []types.NonApprovedLanguage {
	var nonApprovedLanguages []types.NonApprovedLanguage
	for _, lang := range ctx.DownloadedLanguages {
		var vars []types.NewVar
		for name, translation := range lang.LanguageVars {
			t, _ := time.Parse(time.RFC3339, translation.TranslatedAt)
			if !translation.IsApproved {
				vars = append(vars, types.NewVar{
					Name:           name,
					Text:           translation.Text,
					Lang:           lang.LanguageCode,
					TranslationId:  translation.TranslationId,
					TranslatorName: translation.TranslatorName,
					TranslatedAt:   t.Unix(),
				})
			}
		}
		sort.Slice(vars, func(i, j int) bool {
			return vars[i].TranslatedAt > vars[j].TranslatedAt
		})
		if len(vars) > 0 {
			nonApprovedLanguages = append(nonApprovedLanguages, types.NonApprovedLanguage{
				LanguageCode: lang.LanguageCode,
				Vars:         vars,
				LastUpdateAt: vars[0].TranslatedAt,
			})
		}
	}
	sort.Slice(nonApprovedLanguages, func(i, j int) bool {
		return nonApprovedLanguages[i].LastUpdateAt > nonApprovedLanguages[j].LastUpdateAt
	})
	return nonApprovedLanguages
}
