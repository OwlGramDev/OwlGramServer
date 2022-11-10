package types

type UploadFileRsp struct {
	IfSuccess    int        `json:"ifSuccess"`
	FileInfoList []FileInfo `json:"fileInfoList"`
}
