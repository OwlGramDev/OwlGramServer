package types

import "OwlGramServer/compiler/types"

type PythonParams struct {
	Apks     []types.PackageInfo `json:"apks"`
	Base     string              `json:"base"`
	Channel  string              `json:"channel"`
	BotToken string              `json:"bot_token"`
	ApiId    int                 `json:"api_id"`
	ApiHash  string              `json:"api_hash"`
	ChatID   int64               `json:"chat_id"`
}
