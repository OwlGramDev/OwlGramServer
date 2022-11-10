package types

type ReviewInfo struct {
	AuthorName     string  `json:"author_name"`
	AndroidSDK     int8    `json:"android_sdk"`
	AppVersionCode int32   `json:"app_version_code"`
	AppVersionName string  `json:"app_version_name"`
	StarRating     int8    `json:"star_rating"`
	DeviceModel    *string `json:"device_model"`
	CPUVendor      *string `json:"cpu_vendor"`
	Text           string  `json:"text"`
	LastEdit       int64   `json:"last_edit"`
}
