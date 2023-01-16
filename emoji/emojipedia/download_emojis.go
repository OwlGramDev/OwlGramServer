package emojipedia

import (
	"OwlGramServer/emoji/emojipedia/types"
	"OwlGramServer/http"
	"OwlGramServer/utilities/concurrency"
	"bytes"
	"fmt"
	"golang.org/x/image/draw"
	"image"
	"image/png"
	"sync"
)

func downloadEmojis(pd []*types.ProviderDescriptor, emojiResult map[string]*types.EmojiRaw) error {
	var wg sync.WaitGroup
	responses := make(chan *types.EmojiRequest, len(emojiResult)*len(pd))
	semaphore := concurrency.NewPool[string](400)
	for i, data := range emojiResult {
		for _, providerData := range pd {
			emojiLink := fmt.Sprintf("%s%s_%s.png", providerData.Link, data.EmojiName, data.EmojiHex)
			wg.Add(1)
			semaphore.Enqueue(func(params ...string) {
				defer wg.Done()
				link := params[0]
				emoji := params[1]
				provider := params[2]
				emojiPng := http.ExecuteRequest(
					link,
					http.Retries(3),
				)
				if emojiPng.Error == nil && len(emojiPng.Read()) > 1024*500 {
					decode, _ := png.Decode(bytes.NewBuffer(emojiPng.Read()))
					resized := image.NewRGBA(image.Rect(0, 0, 66, 66))
					draw.CatmullRom.Scale(resized, resized.Bounds(), decode, decode.Bounds(), draw.Over, nil)
					var buf bytes.Buffer
					_ = png.Encode(&buf, resized)
					emojiPng.SetFallback(buf.Bytes())
				}
				responses <- &types.EmojiRequest{
					Provider:   provider,
					Emoji:      emoji,
					EmojiLink:  link,
					HttpResult: emojiPng,
				}
			}, emojiLink, i, providerData.ID)
		}
	}
	wg.Wait()
	close(responses)

	for resp := range responses {
		data := resp.HttpResult.Read()
		if len(data) == 0 || resp.HttpResult.Error != nil {
			continue
		}
		for _, providerData := range pd {
			if providerData.ID == resp.Provider {
				providerData.Emojis[resp.Emoji] = data
			}
		}
	}
	return nil
}
