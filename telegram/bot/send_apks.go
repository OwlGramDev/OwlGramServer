package bot

import (
	"OwlGramServer/compiler"
	"OwlGramServer/compiler/types"
	"OwlGramServer/consts"
	"bytes"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"os/exec"
	"strings"
)

func (c *Context) SendApks(chatId int64) error {
	var listApkInfo []types.PackageInfo
	for _, apk := range c.CompilerClient.Bundles {
		if apk.IsApk {
			listApkInfo = append(listApkInfo, types.PackageInfo{
				MinSDKVersion:    apk.MinSDKVersion,
				TargetSDKVersion: apk.TargetSDKVersion,
				VersionCode:      apk.VersionCode,
				VersionName:      apk.VersionName,
				AbiName:          compiler.GetCoolAbi(apk.AbiName, 3),
				Path:             apk.Path,
			})
		}
	}
	marshal, _ := json.Marshal(struct {
		Apks     []types.PackageInfo `json:"apks"`
		Base     string              `json:"base"`
		Channel  string              `json:"channel"`
		BotToken string              `json:"bot_token"`
		ApiId    int                 `json:"api_id"`
		ApiHash  string              `json:"api_hash"`
		ChatID   int64               `json:"chat_id"`
	}{
		Apks:     listApkInfo,
		Base:     c.ReleaseBase,
		Channel:  c.ReleaseType,
		BotToken: consts.BotToken,
		ApiId:    consts.ApiID,
		ApiHash:  consts.ApiHash,
		ChatID:   chatId,
	})
	var errMess bytes.Buffer
	cmd := exec.Command("bash", "-c", strings.Join([]string{
		"source env/bin/activate",
		fmt.Sprintf("python3.10 __main__.py %s", hex.EncodeToString(marshal)),
	}, " && "))
	cmd.Dir = consts.PythonLibApkSenderPath
	cmd.Stderr = &errMess
	_, err := cmd.Output()
	if err != nil {
		if errMess.Len() > 0 {
			return errors.New(errMess.String())
		}
		return err
	}
	return nil
}
