package compiler

import (
	"OwlGramServer/compiler/types"
	"OwlGramServer/consts"
	"OwlGramServer/utilities"
	"fmt"
	"log"
	"os"
	"path"
)

func (context *Context) Run(messageId int64, listener func(int)) {
	context.IsRunning = true
	context.MessageID = messageId
	f, err := os.Open(consts.UploadsFolder)
	if err != nil {
		panic(fmt.Sprintf("Error opening folder %s", consts.UploadsFolder))
		return
	}
	files, err := f.Readdir(0)
	if err != nil {
		panic("Error reading directory")
	}
	zipFile := ""
	for _, v := range files {
		if v.Name()[len(v.Name())-3:] == "zip" && !v.IsDir() {
			zipFile = v.Name()
			break
		}
	}
	if zipFile == "" {
		context.IsRunning = false
		listener(utilities.NoZipFound)
		return
	}
	if _, err := os.Stat(consts.CacheFolderBundles); os.IsNotExist(err) {
		_ = os.Mkdir(consts.CacheFolderBundles, 0775)
	} else {
		_ = os.RemoveAll(consts.CacheFolderBundles)
		_ = os.Mkdir(consts.CacheFolderBundles, 0775)
	}
	if _, err := os.Stat(consts.CacheFolderExtractedApks); os.IsNotExist(err) {
		_ = os.Mkdir(consts.CacheFolderExtractedApks, 0775)
	} else {
		_ = os.RemoveAll(consts.CacheFolderExtractedApks)
		_ = os.Mkdir(consts.CacheFolderExtractedApks, 0775)
	}
	err = UnzipFiles(path.Join(consts.UploadsFolder, zipFile), consts.CacheFolderBundles)
	if err != nil {
		context.IsRunning = false
		log.Println(err)
		listener(utilities.UnzipError)
		return
	}
	listPackages := orderBundles(getFiles(consts.CacheFolderBundles[:len(consts.CacheFolderBundles)-1]))
	if len(listPackages) == 0 {
		context.IsRunning = false
		listener(utilities.NoBundleFound)
		return
	}
	listener(utilities.ReadingBundlesInfo)
	var packagesInfo []types.PackageInfo
	for _, bundle := range listPackages {
		if path.Ext(bundle) == ".apk" {
			bundleInfo, err := GetApkInfo(bundle)
			if err != nil {
				context.IsRunning = false
				listener(utilities.BundleCorrupted)
				return
			}
			packagesInfo = append(packagesInfo, bundleInfo)
		} else {
			bundleInfo, err := GetBundleInfo(bundle)
			if err != nil {
				context.IsRunning = false
				listener(utilities.BundleCorrupted)
				return
			}
			packagesInfo = append(packagesInfo, bundleInfo)
		}
	}
	var versionCode int
	for _, fpk := range packagesInfo {
		if versionCode == 0 {
			versionCode = fpk.VersionCode
		} else if versionCode != fpk.VersionCode {
			context.IsRunning = false
			listener(utilities.BundlesHaveDifferentVersions)
			return
		}
	}
	for index, apk := range packagesInfo {
		r, err := os.ReadFile(apk.Path)
		if err != nil {
			context.IsRunning = false
			listener(utilities.BundleCorrupted)
			return
		}
		apkPath := path.Join(consts.CacheFolderExtractedApks, fmt.Sprintf("%s-SDK%s%s", apk.AbiName, apk.MinSDKVersion, path.Ext(apk.Path)))
		err = os.WriteFile(apkPath, r, 0775)
		packagesInfo[index].Path = apkPath
		if err != nil {
			context.IsRunning = false
			listener(utilities.BundleCorrupted)
			return
		}
		if !apk.IsApk {
			ExtractMapping(apk.Path, apk.VersionCode)
		}
	}
	context.Bundles = packagesInfo
	listener(utilities.BundlesCompiled)
}
