package handlers

import (
	"OwlGramServer/consts"
	"OwlGramServer/telegram/emoji"
	"OwlGramServer/telegram/emoji/types"
	"encoding/json"
	"fmt"
	"github.com/valyala/fasthttp"
)

func EmojiPacks(ctx *fasthttp.RequestCtx, clientEmoji *emoji.Context) {
	if clientEmoji.EmojiPacks == nil {
		Forbidden(ctx)
		return
	}
	ctx.SetContentType("application/json")
	ctx.SetStatusCode(fasthttp.StatusOK)
	var result []types.PacksInfo
	for _, pack := range clientEmoji.EmojiPacks {
		result = append(result, types.PacksInfo{
			Name:     pack.DisplayName,
			Version:  pack.Version,
			Id:       pack.Identifier,
			Preview:  fmt.Sprintf("%s/previews/%s.png?v=%d", consts.ServerBase, pack.Identifier, pack.Version),
			File:     fmt.Sprintf("%s/packs/%s.zip?v=%s", consts.ServerBase, pack.Identifier, pack.MD5),
			FileSize: len(pack.Emojies),
			MD5:      pack.MD5,
		})
	}
	res, _ := json.Marshal(result)
	ctx.SetBody(res)
	ctx.SetConnectionClose()
}
