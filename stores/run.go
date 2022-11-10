package stores

import (
	"OwlGramServer/compiler"
	bundles "OwlGramServer/compiler/types"
	"OwlGramServer/stores/app_gallery"
	"OwlGramServer/stores/github"
	"OwlGramServer/stores/google_play"
	"OwlGramServer/stores/owlgram_server"
	"OwlGramServer/stores/types"
	updates "OwlGramServer/updates/types"
	"OwlGramServer/utilities"
	"strconv"
	"time"
)

func (context *Context) Run(listApks []bundles.PackageInfo, releaseType string, bannerTmp string, hrefTmp string, localizationsTmp map[string]map[string]string, update updates.UpdatesDescriptor, listener func(int)) {

	context.StoreList = []*types.StoreInfo{
		{
			Name:   "OwlGram Server",
			Status: utilities.StatusPending,
		},
		{
			Name:   "GitHub",
			Status: utilities.StatusPending,
		},
		{
			Name:   "Google Play",
			Status: utilities.StatusPending,
		},
		{
			Name:   "AppGallery",
			Status: utilities.StatusPending,
		},
	}
	listener(utilities.SendingToStores)
	updateDescription := context.updateClient.GetChangelogUpdate(strconv.Itoa(listApks[0].VersionCode), "en", true)
	for index := range context.StoreList {
		context.StoreList[index].Status = utilities.StatusUploading
		listener(utilities.SendingToStores)
		handleStatus := func(progress float64) {
			if progress == -1 {
				context.StoreList[index].Status = utilities.StatusFailed
				for i := index + 1; i < len(context.StoreList); i++ {
					context.StoreList[i].Status = utilities.StatusCanceled
				}
				listener(utilities.FailedToSendToStores)
			} else if progress == -2 {
				context.StoreList[index].Status = utilities.StatusSuccess
				listener(utilities.SendingToStores)
			} else if progress == -3 {
				context.StoreList[index].Status = utilities.StatusElaborating
				listener(utilities.SendingToStores)
			} else {
				context.StoreList[index].Progress = progress
				listener(utilities.SendingToStores)
			}
		}
		if context.StoreList[index].Status == utilities.StatusCanceled {
			return
		}
		switch context.StoreList[index].Name {
		case "GitHub":
			github.Upload(listApks, releaseType, *updateDescription, handleStatus)
			break
		case "Google Play":
			google_play.Upload(listApks, handleStatus)
			break
		case "AppGallery":
			universal := compiler.GetUniversal(listApks)
			if universal == nil {
				handleStatus(-1)
				return
			}
			app_gallery.Upload(*universal, *updateDescription, handleStatus)
			break
		case "OwlGram Server":
			owlgram_server.Upload(listApks, bannerTmp, hrefTmp, localizationsTmp, releaseType, update, handleStatus)
			break
		}
		time.Sleep(time.Millisecond * 500)
	}
}
