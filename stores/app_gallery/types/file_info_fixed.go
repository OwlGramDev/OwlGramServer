package types

type FileInfoFixed struct {
	FileName    string `json:"fileName"`
	FileDestUrl string `json:"fileDestUrl"`
	Size        int64  `json:"size"`
}
