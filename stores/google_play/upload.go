package google_play

import (
	bundles "OwlGramServer/compiler/types"
	"OwlGramServer/consts"
	"bufio"
	"context"
	"google.golang.org/api/androidpublisher/v3"
	"google.golang.org/api/googleapi"
	"google.golang.org/api/option"
	"math"
	"os"
	"strconv"
	"time"
)

func Upload(bundleList []bundles.PackageInfo, listener func(float64)) {
	var listApksTmp []bundles.PackageInfo
	for _, apk := range bundleList {
		if !apk.IsApk {
			listApksTmp = append(listApksTmp, apk)
		}
	}
	bundleList = listApksTmp
	ctx := context.Background()
	pubService, err := androidpublisher.NewService(ctx, option.WithCredentialsFile(consts.GoogleApiJsonPath))
	if err != nil {
		listener(-1)
		return
	}
	appEdit, err := pubService.Edits.Insert(consts.AndroidPackageName, &androidpublisher.AppEdit{
		ExpiryTimeSeconds: strconv.Itoa(10 * 60),
	}).Do()
	if err != nil {
		listener(-1)
		return
	}
	totalProgress := make([]int, len(bundleList))
	foundError := false
	waitUntilComplete := make(chan bool)
	for i := 0; i < len(bundleList); i++ {
		go func(i int) {
			f, err := os.Open(bundleList[i].Path)
			if err != nil {
				foundError = true
				return
			}
			r := bufio.NewReader(f)
			totalSize, _ := f.Stat()
			uploadRequest := pubService.Edits.Bundles.Upload(consts.AndroidPackageName, appEdit.Id)
			uploadRequest.Media(r, googleapi.ContentType("application/octet-stream"))
			uploadRequest.ProgressUpdater(func(current, total int64) {
				totalProgress[i] = int(math.Round(float64(current) / float64(totalSize.Size()) * 100))
			})
			go func() {
				for {
					if foundError {
						_, _ = r.Discard(1)
					}
					time.Sleep(time.Millisecond * 100)
				}
			}()
			_, err = uploadRequest.Do()
			if err != nil {
				foundError = true
				return
			}
		}(i)
	}
	go func() {
		for {
			progress := 0
			for i := 0; i < len(totalProgress); i++ {
				progress += totalProgress[i]
			}
			progress = progress / len(totalProgress)
			listener(float64(progress))
			if progress == 100 || foundError {
				waitUntilComplete <- foundError
				break
			}
			time.Sleep(time.Second)
		}
	}()
	if <-waitUntilComplete {
		listener(-1)
		return
	}
	_, err = pubService.Edits.Commit(consts.AndroidPackageName, appEdit.Id).Do()
	if err != nil {
		listener(-1)
		return
	}
	listener(-2)
}
