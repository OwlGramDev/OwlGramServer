package app_gallery

import (
	"OwlGramServer/compiler"
	bundles "OwlGramServer/compiler/types"
	"OwlGramServer/consts"
	"OwlGramServer/http"
	multipart "OwlGramServer/http/types"
	"OwlGramServer/stores/app_gallery/types"
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"strings"
	"time"
)

func Upload(apk bundles.PackageInfo, descriptionUpdate string, listener func(float64)) {
	re := regexp.MustCompile(`<[^>]*>?`)
	descriptionUpdate = re.ReplaceAllString(descriptionUpdate, "")
	descriptionUpdate = strings.Replace(descriptionUpdate, strings.Split(descriptionUpdate, "\n\n")[0]+"\n\n", "", -1)
	descriptionUpdate = consts.EmojiRx.ReplaceAllString(descriptionUpdate, "")
	if len(descriptionUpdate) > 1000 {
		descriptionUpdate = descriptionUpdate[:996] + "..."
	}
	marshal, _ := json.Marshal(types.OAuth{
		ClientId:     consts.HuaweiClientId,
		ClientSecret: consts.HuaweiClientSecret,
		GrantType:    "client_credentials",
	})
	res := http.ExecuteRequest(
		fmt.Sprintf("%soauth2/v1/token", consts.AppGalleryApi),
		http.Method("POST"),
		http.Headers(map[string]string{
			"Content-Type": "application/json",
		}),
		http.Body(marshal),
	)
	if res.Error != nil {
		fmt.Println(res.Error)
		listener(-1)
		return
	}
	var authToken types.AuthToken
	err := json.Unmarshal(res.Read(), &authToken)
	if err != nil {
		listener(-1)
		return
	}
	if len(descriptionUpdate) > 0 {
		marshal, _ = json.Marshal(types.AppLanguageInfo{
			Lang:        "en-US",
			NewFeatures: descriptionUpdate,
		})
		res = http.ExecuteRequest(
			fmt.Sprintf("%spublish/v2/app-language-info?appId=%d", consts.AppGalleryApi, consts.AppGalleryAppID),
			http.Method("PUT"),
			http.BearerToken(authToken.AccessToken),
			http.Headers(map[string]string{
				"client_id":    consts.HuaweiClientId,
				"Content-Type": "application/json",
			}),
			http.Body(marshal),
		)
		if res.Error != nil {
			listener(-1)
			return
		}
	}
	res = http.ExecuteRequest(
		fmt.Sprintf("%spublish/v2/upload-url?appId=%d&suffix=apk", consts.AppGalleryApi, consts.AppGalleryAppID),
		http.Method("GET"),
		http.BearerToken(authToken.AccessToken),
		http.Headers(map[string]string{
			"client_id": consts.HuaweiClientId,
		}),
	)
	if res.Error != nil {
		listener(-1)
		return
	}
	var uploadUrl types.UploadUrl
	err = json.Unmarshal(res.Read(), &uploadUrl)
	if err != nil {
		listener(-1)
		return
	}
	input, err := os.ReadFile(apk.Path)
	if err != nil {
		listener(-1)
		return
	}
	res = http.ExecuteRequest(
		uploadUrl.UploadUrl,
		http.Method("POST"),
		http.MultiPartForm(
			http.Data(map[string]string{
				"authCode":  uploadUrl.AuthCode,
				"fileCount": "1",
			}),
			http.Files(map[string]multipart.FileDescriptor{
				"file": {
					FileName: "universal.apk",
					Content:  input,
				},
			}),
		),
	)
	if res.Error != nil {
		listener(-1)
		return
	}
	listener(100)
	var uploadResult types.UploadResult
	err = json.Unmarshal(res.Read(), &uploadResult)
	if err != nil {
		listener(-1)
		return
	}
	var fixedFileList []types.FileInfoFixed
	for _, file := range uploadResult.Result.UploadFileRsp.FileInfoList {
		fixedFileList = append(fixedFileList, types.FileInfoFixed{
			FileName:    fmt.Sprintf("OwlGram %s %s.apk", apk.VersionName, compiler.GetCoolAbi(apk.AbiName, 0)),
			FileDestUrl: file.FileDestUlr,
			Size:        file.Size,
		})
	}
	appFileInfo := types.AppFileInfo{
		Files:    fixedFileList,
		FileType: 5,
		Lang:     "en-US",
	}
	marshal, _ = json.Marshal(appFileInfo)
	res = http.ExecuteRequest(
		fmt.Sprintf("%spublish/v2/app-file-info?appId=%d", consts.AppGalleryApi, consts.AppGalleryAppID),
		http.Method("PUT"),
		http.BearerToken(authToken.AccessToken),
		http.Headers(map[string]string{
			"Content-Type": "application/json",
			"client_id":    consts.HuaweiClientId,
		}),
		http.Body(marshal),
	)
	if res.Error != nil {
		listener(-1)
		return
	}
	sentCorrectly := false
	listener(-3)
	for {
		res = http.ExecuteRequest(
			fmt.Sprintf("%spublish/v2/app-submit?appid=%d", consts.AppGalleryApi, consts.AppGalleryAppID),
			http.Method("POST"),
			http.BearerToken(authToken.AccessToken),
			http.Headers(map[string]string{
				"Content-Type": "application/json",
				"client_id":    consts.HuaweiClientId,
			}),
		)
		if res.Error != nil {
			return
		}
		var submitResult types.SubmitResult
		_ = json.Unmarshal(res.Read(), &submitResult)
		if submitResult.Ret.Code == 204144727 {
			time.Sleep(time.Second * 1)
		} else if submitResult.Ret.Code == 0 {
			sentCorrectly = true
			break
		} else {
			sentCorrectly = false
			break
		}
	}
	if sentCorrectly {
		listener(-2)
	} else {
		listener(-1)
	}
}
