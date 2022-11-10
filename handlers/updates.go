package handlers

import (
	"OwlGramServer/consts"
	"OwlGramServer/updates"
	"OwlGramServer/updates/types"
	"encoding/json"
	"fmt"
	"github.com/valyala/fasthttp"
	"os"
	"path"
	"strconv"
)

func Updates(ctx *fasthttp.RequestCtx, clientUpdates *updates.Context) {
	lang := ctx.Request.URI().QueryArgs().Peek("lang")
	isBetaByte := ctx.Request.URI().QueryArgs().Peek("beta")
	abiByte := ctx.Request.URI().QueryArgs().Peek("abi")
	if lang != nil && isBetaByte != nil && abiByte != nil {
		isBeta := string(isBetaByte) == "true"
		abi := string(abiByte)
		ctx.SetContentType("application/json")
		ctx.SetStatusCode(fasthttp.StatusOK)
		updatesInfo := clientUpdates.UpdatesDescriptor
		if updatesInfo == nil {
			Forbidden(ctx)
			return
		}
		var updateInfo *types.UpdateInfo
		if isBeta {
			isBeta = updatesInfo.Updates.Beta.VersionCode >= updatesInfo.Updates.Stable.VersionCode
		}
		isBeta = isBeta && updatesInfo.Updates.Stable.VersionCode != updatesInfo.Updates.Beta.VersionCode
		channelName := "stable"
		if isBeta {
			channelName = "beta"
			updateInfo = updatesInfo.Updates.Beta
		} else {
			updateInfo = updatesInfo.Updates.Stable
		}
		localization := updates.GetLocalization(string(lang), updatesInfo.Localizations)
		isAvailableDownload := false
		fileName := updates.GetFile(isBeta, abi)
		fileFolder := path.Join(consts.ServerFilesFolder, "versions", strconv.Itoa(updateInfo.VersionCode), fileName)
		if _, err := os.Stat(fileFolder); err == nil {
			isAvailableDownload = true
		}
		if updatesInfo.Updates.Beta.VersionCode == 0 || updatesInfo.Updates.Stable.VersionCode == 0 {
			isAvailableDownload = false
		}
		if isAvailableDownload {
			fi, err := os.Stat(fileFolder)
			if err != nil {
				return
			}
			size := fi.Size()
			res, _ := json.Marshal(
				types.Update{
					Status:   "update_available",
					Title:    fmt.Sprintf("<b>%s</b>", fmt.Sprintf(localization["title"], updateInfo.VersionName)),
					Desc:     fmt.Sprintf("<b>%s</b>\n<a href=\"%s\"><u><b>%s</b></u></a>", localization["desc1"], updateInfo.Href, localization[fmt.Sprintf("desc2_%s", channelName)]),
					Note:     localization[fmt.Sprintf("note_%s", channelName)],
					Banner:   updateInfo.Banner,
					LinkFile: fmt.Sprintf("%sversions/%d/%s", consts.OwlGramFilesServer, updateInfo.VersionCode, fileName),
					FileSize: size,
					Version:  updateInfo.VersionCode,
				},
			)
			ctx.SetBody(res)
		} else {
			res, _ := json.Marshal(types.NoUpdate{
				Status: "no_updates",
			})
			ctx.SetBody(res)
		}
		ctx.SetConnectionClose()
		return
	}
	Forbidden(ctx)
}
