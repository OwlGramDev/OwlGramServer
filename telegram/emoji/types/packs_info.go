package types

type PacksInfo struct {
	Name     string `json:"name"`
	Version  int    `json:"version"`
	Id       string `json:"id"`
	Preview  string `json:"preview"`
	File     string `json:"file"`
	FileSize int    `json:"file_size"`
}
