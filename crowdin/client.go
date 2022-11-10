package crowdin

func Client() *Context {
	return &Context{
		LanguagesList:       nil,
		DownloadedLanguages: nil,
	}
}
