package crowdin

import (
	"strings"
)

func (ctx *Context) GetLanguageVars(code string) *map[string]string {
	if len(code) == 0 {
		return nil
	}
	for i := 0; i < len(ctx.DownloadedLanguages); i++ {
		if strings.HasPrefix(ctx.DownloadedLanguages[i].LanguageCode, code) {
			listVars := make(map[string]string)
			for key, value := range ctx.DownloadedLanguages[i].LanguageVars {
				if value.IsApproved {
					listVars[key] = value.Text
				} else {
					listVars[key] = value.ApprovedText
				}
			}
			return &listVars
		}
	}
	if strings.Contains(code, "-") {
		return ctx.GetLanguageVars(strings.Split(code, "-")[0])
	}
	return nil
}
