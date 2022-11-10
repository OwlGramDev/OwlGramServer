package reviews

import (
	"OwlGramServer/google_reviews/types"
	"crypto/sha256"
	"encoding/hex"
	"strconv"
)

func (ctx *Context) getMD5Review(review types.ReviewInfo) string {
	reviewGenerator := []byte(review.Text)
	reviewGenerator = append(reviewGenerator, []byte(review.AuthorName)...)
	reviewGenerator = append(reviewGenerator, []byte(strconv.Itoa(int(review.StarRating)))...)
	reviewGenerator = append(reviewGenerator, []byte(strconv.Itoa(int(review.LastEdit)))...)
	byteSum := sha256.Sum256(reviewGenerator)
	hash256 := hex.EncodeToString(byteSum[:])
	return hash256
}
