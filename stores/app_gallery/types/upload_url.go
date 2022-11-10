package types

type UploadUrl struct {
	UploadUrl      string `json:"uploadUrl"`
	ChunkUploadUrl string `json:"chunkUploadUrl"`
	AuthCode       string `json:"authCode"`
}
