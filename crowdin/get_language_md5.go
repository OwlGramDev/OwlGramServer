package crowdin

import "strings"

func (ctx *Context) GetLanguageMD5(code string) *string {
	if len(code) == 0 {
		return nil
	}
	for i := 0; i < len(ctx.DownloadedLanguages); i++ {
		if strings.HasPrefix(ctx.DownloadedLanguages[i].LanguageCode, code) {
			return &ctx.DownloadedLanguages[i].MD5
		}
	}
	if strings.Contains(code, "-") {
		return ctx.GetLanguageMD5(strings.Split(code, "-")[0])
	}
	return nil
}
