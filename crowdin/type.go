package crowdin

import (
	"OwlGramServer/crowdin/types"
)

type Context struct {
	SourceStrings       map[int]string
	LanguagesList       []types.LanguageData
	DownloadedLanguages []types.Language
}
