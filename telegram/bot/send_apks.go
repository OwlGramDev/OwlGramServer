package bot

import (
	"OwlGramServer/compiler"
	"OwlGramServer/compiler/types"
	"OwlGramServer/consts"
	typesBot "OwlGramServer/telegram/bot/types"
	"path"
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
	_, err := c.pythonClient.Run(
		path.Join(consts.PythonLibApkSenderPath, "__main__.py"),
		typesBot.PythonParams{
			Apks:     listApkInfo,
			Base:     c.ReleaseBase,
			Channel:  c.ReleaseType,
			BotToken: consts.BotToken,
			ApiId:    consts.ApiID,
			ApiHash:  consts.ApiHash,
			ChatID:   chatId,
		},
	)
	return err
}
