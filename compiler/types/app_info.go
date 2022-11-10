package types

type PackageInfo struct {
	AbiName          string `json:"abi_name"`
	MinSDKVersion    string `json:"min_sdk_version"`
	TargetSDKVersion string `json:"target_sdk_version"`
	VersionCode      int    `json:"version_code"`
	VersionName      string `json:"version_name"`
	Path             string `json:"path"`
	IsApk            bool   `json:"-"`
}
