package translator

import "OwlGramServer/translator/types"

func Translate(query, destLanguage string) *types.Result {
	blocks := getStringBlocks(query, 2500)
	apiUrl := "https://translate.google.com/translate_a/single?dj=1" +
		"&q=%s" +
		"&sl=auto" +
		"&tl=%s" +
		"&ie=UTF-8&oe=UTF-8&client=at&dt=t&otf=2"
	var resultString string
	var resultLanguage string
	for _, block := range blocks {
		result := translateBlock(apiUrl, block, destLanguage)
		if result == nil {
			return nil
		}
		if resultLanguage == "" {
			resultLanguage = result.SourceLanguage
		}
		resultString += buildTranslatedString(block, result.Translation)
	}
	return &types.Result{
		Translation:    resultString,
		SourceLanguage: resultLanguage,
	}
}
