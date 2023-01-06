package types

import (
	"crypto/sha256"
	"encoding/hex"
	"strconv"
)

type ReviewInfo struct {
	ID             string  `json:"id"`
	AuthorName     string  `json:"author_name"`
	AndroidSDK     int8    `json:"android_sdk"`
	AppVersionCode int32   `json:"app_version_code"`
	AppVersionName string  `json:"app_version_name"`
	StarRating     int8    `json:"star_rating"`
	DeviceModel    *string `json:"device_model"`
	CPUVendor      *string `json:"cpu_vendor"`
	Text           string  `json:"text"`
	LastEdit       int64   `json:"last_edit"`
}

func (r *ReviewInfo) GetMD5() string {
	reviewGenerator := []byte(r.Text)
	reviewGenerator = append(reviewGenerator, []byte(strconv.Itoa(int(r.StarRating)))...)
	reviewGenerator = append(reviewGenerator, []byte(strconv.Itoa(int(r.LastEdit)))...)
	byteSum := sha256.Sum256(reviewGenerator)
	hash256 := hex.EncodeToString(byteSum[:])
	return r.ID + hash256
}
