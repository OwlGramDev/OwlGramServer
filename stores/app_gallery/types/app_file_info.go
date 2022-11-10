package types

type AppFileInfo struct {
	Files    []FileInfoFixed `json:"files"`
	FileType int             `json:"fileType"`
	Lang     string          `json:"lang"`
}
