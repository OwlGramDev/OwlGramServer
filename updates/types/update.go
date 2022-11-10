package types

type Update struct {
	Status   string `json:"status"`
	Title    string `json:"title"`
	Desc     string `json:"desc"`
	Note     string `json:"note"`
	Banner   string `json:"banner"`
	LinkFile string `json:"link_file"`
	FileSize int64  `json:"file_size"`
	Version  int    `json:"version"`
}
