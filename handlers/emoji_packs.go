package handlers

import (
	"OwlGramServer/consts"
	"OwlGramServer/emoji"
	"OwlGramServer/emoji/types"
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
			Name:       pack.Name,
			Version:    pack.UnicodeVersion,
			ID:         emoji.LegacyID(pack.ID, pack.GetID()),
			Preview:    fmt.Sprintf("%s/previews/%s.png?v=%s", consts.ServerBase, pack.ID, pack.MD5),
			File:       fmt.Sprintf("%s/packs/%s.zip?v=%s", consts.ServerBase, pack.ID, pack.MD5),
			FileSize:   len(pack.EmojiZip),
			EmojiCount: pack.EmojiCount,
			MD5:        pack.MD5,
		})
	}
	res, _ := json.Marshal(result)
	ctx.SetBody(res)
	ctx.SetConnectionClose()
}
