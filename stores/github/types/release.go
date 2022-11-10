package types

type Release struct {
	TagName    string `json:"tag_name"`
	Name       string `json:"name"`
	Body       string `json:"body"`
	PreRelease bool   `json:"prerelease"`
}
