package translator

import (
	"OwlGramServer/http"
	"OwlGramServer/translator/types"
	"fmt"
	"math/rand"
	"net/url"
)

func translateBlock(apiUrl, block, destLanguage string) *types.Result {
	request := http.ExecuteRequest(fmt.Sprintf(apiUrl, url.QueryEscape(block), destLanguage), http.Headers(map[string]string{
		"User-Agent": fmt.Sprintf("GoogleTranslate/6.28.0.05.421483610 (%s)", devices[rand.Intn(len(devices))]),
	}))
	if request.Error != nil {
		return nil
	}
	return getResult(request.Read())
}
