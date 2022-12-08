package translator

import (
	"OwlGramServer/translator/types"
	"encoding/json"
)

func getResult(request []byte) *types.Result {
	var response types.GoogleResult
	var stringBuilder string
	err := json.Unmarshal(request, &response)
	if err != nil {
		return nil
	}
	for _, sentence := range response.Sentences {
		stringBuilder += sentence.Trans
	}
	return &types.Result{
		Translation:    stringBuilder,
		SourceLanguage: response.Src,
	}
}
