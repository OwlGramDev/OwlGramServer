package types

type FileInfo struct {
	FileDestUlr     string `json:"fileDestUlr"`
	Size            int64  `json:"size"`
	DisposableURL   string `json:"disposableUrl"`
	PurifiedForFile bool   `json:"purifiedForFile"`
}
