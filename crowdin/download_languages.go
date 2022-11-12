package crowdin

import (
	"OwlGramServer/consts"
	"OwlGramServer/crowdin/types"
	"OwlGramServer/http"
	"encoding/json"
	"fmt"
	"log"
	"sync"
)

func (ctx *Context) downloadLanguages() {
	listLanguagesRaw, _ := http.ExecuteRequest(
		fmt.Sprintf("%s%s", consts.CrowdinApiLink, consts.CrowdinProjectId),
		http.BearerToken(consts.CrowdinAuthToken),
		http.Retries(3),
	)
	var languageIdentifiers struct {
		Data []struct {
			Data struct {
				Id         int    `json:"id"`
				Identifier string `json:"identifier"`
			} `json:"data"`
		} `json:"data"`
	}
	r, _ := http.ExecuteRequest(
		fmt.Sprintf("%s%s/strings?limit=500", consts.CrowdinApiLink, consts.CrowdinProjectId),
		http.BearerToken(consts.CrowdinAuthToken),
		http.Retries(3),
	)
	err := json.Unmarshal(r, &languageIdentifiers)
	if err != nil {
		return
	}
	if listLanguagesRaw != nil && err == nil && r != nil {
		var projectMain types.ProjectInfo
		_ = json.Unmarshal(listLanguagesRaw, &projectMain)
		ctx.LanguagesList = projectMain.Data.TargetLanguages
		ctx.SourceStrings = make(map[int]string)
		for _, language := range languageIdentifiers.Data {
			ctx.SourceStrings[language.Data.Id] = language.Data.Identifier
		}
		var listLanguages []types.Language
		translations, _ := ctx.GetTranslationsExported("en", false)
		languageElementBase := types.Language{
			MD5:          translations.MD5,
			LanguageCode: "en",
			LanguageVars: make(map[string]types.TranslationInfo),
		}
		for key, value := range translations.Translations {
			languageElementBase.LanguageVars[key] = types.TranslationInfo{
				Text:       value,
				IsApproved: true,
			}
		}
		listLanguages = append(listLanguages, languageElementBase)
		var waitUntilFinish sync.WaitGroup
		if len(ctx.LanguagesList) > 0 {
			for i := 0; i < len(ctx.LanguagesList); i++ {
				waitUntilFinish.Add(1)
				go func(i int) {
					defer waitUntilFinish.Done()
					translationsOnlyApproved, err1 := ctx.GetTranslationsExported(ctx.LanguagesList[i].Id, true)
					translations, err2 := ctx.GetTranslationsExported(ctx.LanguagesList[i].Id, false)
					translationsInfo, err3 := ctx.GetTranslationInfo(ctx.LanguagesList[i].Id)
					if err1 == nil && err2 == nil && err3 == nil {
						var languageElement types.Language
						languageElement.LanguageCode = ctx.LanguagesList[i].AndroidCode
						languageElement.LanguageVars = make(map[string]types.TranslationInfo)
						languageElement.MD5 = translations.MD5 + translationsOnlyApproved.MD5
						for key, value := range translations.Translations {
							lastApprovedTranslation := translationsOnlyApproved.Translations[key]
							if lastApprovedTranslation == "" {
								lastApprovedTranslation = languageElementBase.LanguageVars[key].Text
							}
							languageElement.LanguageVars[key] = types.TranslationInfo{
								Id:             translationsInfo[key].Id,
								TranslationId:  translationsInfo[key].TranslationId,
								Text:           value,
								ApprovedText:   lastApprovedTranslation,
								IsApproved:     translationsOnlyApproved.Translations[key] != "" && translationsOnlyApproved.Translations[key] == value,
								TranslatorName: translationsInfo[key].User.FullName,
								TranslatedAt:   translationsInfo[key].CreatedAt,
							}
						}
						listLanguages = append(listLanguages, languageElement)
					} else {
						log.Println(err1, err2, err3)
					}
				}(i)
			}
			waitUntilFinish.Wait()
			ctx.DownloadedLanguages = listLanguages
		}
	} else {
		log.Println("Error:", "Can't get languages list")
	}
}
