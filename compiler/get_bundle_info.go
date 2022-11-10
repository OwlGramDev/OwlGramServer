package compiler

import (
	"OwlGramServer/compiler/types"
	"OwlGramServer/consts"
	"os/exec"
	"regexp"
	"strconv"
)

func GetBundleInfo(bundlePath string) (types.PackageInfo, error) {
	cmd := exec.Command("java", "-jar", consts.BundleToolPath, "dump", "manifest", "--bundle", bundlePath)
	output, err := cmd.Output()
	if err != nil {
		return types.PackageInfo{}, err
	}
	r, _ := regexp.Compile(`android:minSdkVersion="(.*?)"`)
	minSdkVersion := r.FindStringSubmatch(string(output))[1]
	r, _ = regexp.Compile(`android:targetSdkVersion="(.*?)"`)
	targetSdkVersion := r.FindStringSubmatch(string(output))[1]
	r, _ = regexp.Compile(`android:versionCode="(.*?)"`)
	versionCodeTmp := r.FindStringSubmatch(string(output))[1]
	versionCode, _ := strconv.Atoi(versionCodeTmp)
	abiCode := versionCode % 10
	versionCode = versionCode / 10
	r, _ = regexp.Compile(`android:versionName="(.*?)"`)
	versionName := r.FindStringSubmatch(string(output))[1]
	return types.PackageInfo{
		AbiName:          GetAbiName(abiCode),
		MinSDKVersion:    minSdkVersion,
		TargetSDKVersion: targetSdkVersion,
		VersionCode:      versionCode,
		VersionName:      versionName,
		Path:             bundlePath,
		IsApk:            false,
	}, nil
}
