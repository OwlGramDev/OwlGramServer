package tg_bot

import (
	"OwlGramServer/consts"
	"OwlGramServer/tg_bot/types"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
)

func CheckTelegramAuthorization(secretData []byte) bool {
	var result types.WebAppInit
	_ = json.Unmarshal(secretData, &result)
	checkString := fmt.Sprintf("auth_date=%s\n", result.AuthDate)
	checkString += fmt.Sprintf("query_id=%s\n", result.QueryId)
	marshal, _ := json.Marshal(result.User)
	checkString += fmt.Sprintf("user=%s", marshal)
	secret := hmac.New(sha256.New, []byte("WebAppData"))
	secret.Write([]byte(consts.BotToken))
	signature := hmac.New(sha256.New, secret.Sum(nil))
	signature.Write([]byte(checkString))
	return hex.EncodeToString(signature.Sum(nil)) == result.Hash
}
