package types

type UploadResult struct {
	Result struct {
		UploadFileRsp UploadFileRsp `json:"uploadFileRsp"`
		ResultCode    string        `json:"resultCode"`
	} `json:"result"`
}
