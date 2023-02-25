package owlgram_server

import (
	"OwlGramServer/compiler"
	bundles "OwlGramServer/compiler/types"
	"OwlGramServer/consts"
	"OwlGramServer/updates/types"
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"path"
	"strconv"
	"strings"
)

func Upload(bundleList []bundles.PackageInfo, bannerTmp string, hrefTmp string, localizationsTmp map[string]map[string]string, releaseType string, update types.UpdatesDescriptor, listener func(float64)) {
	var listApksTmp []bundles.PackageInfo
	for _, apk := range bundleList {
		if apk.IsApk {
			listApksTmp = append(listApksTmp, apk)
		}
	}
	bundleList = listApksTmp
	totalCount := float64(len(bundleList))
	if releaseType == "stable" {
		totalCount += 1
	}
	for i, bundle := range bundleList {
		r, err := os.ReadFile(bundle.Path)
		if err != nil {
			listener(-1)
			return
		}
		pathSave := consts.ServerFilesFolder
		if consts.IsDebug {
			pathSave = consts.DebugFilesFolder
		}
		pathSave = path.Join(pathSave, "versions", strconv.Itoa(bundle.VersionCode))
		if _, err := os.Stat(pathSave); os.IsNotExist(err) {
			_ = os.Mkdir(pathSave, 0755)
		}
		err = os.WriteFile(path.Join(pathSave, fmt.Sprintf("%s-%s.apk", compiler.GetCoolAbi(bundle.AbiName, 1), releaseType)), r, 0755)
		if err != nil {
			listener(-1)
			return
		}
		listener((float64(i) / totalCount) * 100.0)
	}
	if releaseType == "stable" {
		foundUniversal := false
		for _, bundle := range bundleList {
			if bundle.AbiName == "universal" && bundle.IsApk {
				foundUniversal = true
				r, err := os.ReadFile(bundle.Path)
				if err != nil {
					listener(-1)
					return
				}
				pathSave := consts.ServerFilesFolder
				if consts.IsDebug {
					pathSave = consts.DebugFilesFolder
				}
				err = os.WriteFile(path.Join(pathSave, "universal-stable.apk"), r, 0755)
				if err != nil {
					listener(-1)
					return
				}
				listener(100.0)
			}
		}
		if !foundUniversal {
			listener(-1)
			return
		}
	}
	updateFileNew := types.UpdatesDescriptor{
		Updates:       &types.UpdateList{},
		Localizations: localizationsTmp,
	}
	if path.Dir(bannerTmp) != path.Dir(consts.ServerFilesFolder) && !strings.Contains(bannerTmp, "https:") {
		r, err := os.ReadFile(bannerTmp)
		if err != nil {
			listener(-1)
			return
		}
		bannerTmp = path.Join(consts.ServerFilesFolder, fmt.Sprintf("%s.png", RandName()))
		err = os.WriteFile(bannerTmp, r, 0775)
		if err != nil {
			listener(-1)
			return
		}
	}
	bannerTmp = fmt.Sprintf("%s%s", consts.OwlGramFilesServer, path.Base(bannerTmp))

	if releaseType == "stable" {
		updateFileNew.Updates.Beta = update.Updates.Beta
		updateFileNew.Updates.Stable = &types.UpdateInfo{
			Banner: bannerTmp,
			Href:   hrefTmp,
		}
	} else {
		updateFileNew.Updates.Stable = update.Updates.Stable
		updateFileNew.Updates.Beta = &types.UpdateInfo{
			Banner: bannerTmp,
			Href:   hrefTmp,
		}
	}
	marshal, err := json.Marshal(updateFileNew)
	if err != nil {
		listener(-1)
		return
	}
	var out bytes.Buffer
	err = json.Indent(&out, marshal, "", "  ")
	if err != nil {
		listener(-1)
		return
	}
	err = os.WriteFile(consts.UpdateFileDescription, out.Bytes(), 0775)
	if err != nil {
		listener(-1)
		return
	}
	listener(-2)
}
