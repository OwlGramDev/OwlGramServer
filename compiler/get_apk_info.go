package compiler

import (
	"OwlGramServer/compiler/types"
	"OwlGramServer/consts"
	"os/exec"
	"regexp"
	"strconv"
)

func GetApkInfo(path string) (types.PackageInfo, error) {
	cmd := exec.Command(consts.AAPT2ToolPath, "dump", "badging", path)
	output, err := cmd.Output()
	if err != nil {
		return types.PackageInfo{}, err
	}
	r, _ := regexp.Compile(`sdkVersion:'(.*?)'`)
	minSdkVersion := r.FindStringSubmatch(string(output))[1]
	r, _ = regexp.Compile(`targetSdkVersion:'(.*?)'`)
	targetSdkVersion := r.FindStringSubmatch(string(output))[1]
	r, _ = regexp.Compile(`versionCode='(.*?)'`)
	versionCodeTmp := r.FindStringSubmatch(string(output))[1]
	versionCode, _ := strconv.Atoi(versionCodeTmp)
	abiCode := versionCode % 10
	versionCode = versionCode / 10
	r, _ = regexp.Compile(`versionName='(.*?)'`)
	versionName := r.FindStringSubmatch(string(output))[1]
	return types.PackageInfo{
		AbiName:          GetAbiName(abiCode),
		MinSDKVersion:    minSdkVersion,
		TargetSDKVersion: targetSdkVersion,
		VersionCode:      versionCode,
		VersionName:      versionName,
		Path:             path,
		IsApk:            true,
	}, nil
}
