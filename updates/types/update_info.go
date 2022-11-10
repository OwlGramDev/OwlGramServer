package types

type UpdateInfo struct {
	VersionName string `json:"-"`
	VersionCode int    `json:"-"`
	Banner      string `json:"banner"`
	Href        string `json:"href"`
}
