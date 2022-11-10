package google_reviews

import (
	"OwlGramServer/consts"
	"OwlGramServer/google_reviews/types"
	"fmt"
	"regexp"
)

func (ctx *Context) GetAllReview() *[]types.ReviewInfo {
	resultReviews, err := ctx.PlayServiceAPI_.Reviews.List(consts.AndroidPackageName).Do()
	if err != nil {
		return nil
	}
	var listReviews []types.ReviewInfo
	for element := range resultReviews.Reviews {
		reviewInfo := resultReviews.Reviews[element]
		for commentIndex := range reviewInfo.Comments {
			comment := reviewInfo.Comments[commentIndex].UserComment
			if comment != nil {
				var DeviceModel *string
				var CPUVendor *string
				if comment.DeviceMetadata != nil {
					DeviceModel = &comment.DeviceMetadata.ProductName
					rawCPUName := fmt.Sprintf("%s %s", comment.DeviceMetadata.CpuMake, comment.DeviceMetadata.CpuModel)
					CPUVendor = &rawCPUName
					r, _ := regexp.Compile(`.*? \((.*?)\)`)
					result := r.FindAllSubmatch([]byte(*DeviceModel), -1)
					if len(result) > 0 {
						matchResult := string(result[0][1])
						DeviceModel = &matchResult
					}
				}
				var LastModified int64
				if comment.LastModified != nil {
					LastModified = comment.LastModified.Seconds
				}
				listReviews = append(listReviews, types.ReviewInfo{
					AuthorName:     reviewInfo.AuthorName,
					AndroidSDK:     int8(comment.AndroidOsVersion),
					AppVersionCode: int32(comment.AppVersionCode),
					AppVersionName: comment.AppVersionName,
					StarRating:     int8(comment.StarRating),
					DeviceModel:    DeviceModel,
					CPUVendor:      CPUVendor,
					Text:           comment.Text,
					LastEdit:       LastModified,
				})
			}
		}
	}
	return &listReviews
}
