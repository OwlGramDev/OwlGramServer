package emojipedia

import (
	"OwlGramServer/consts"
	"OwlGramServer/emoji/emojipedia/types"
	"crypto/sha256"
	"encoding/hex"

	typesScheme "OwlGramServer/emoji/scheme/types"
	"OwlGramServer/gopy"
	"archive/zip"
	"bytes"
	"encoding/json"
	"fmt"
	"golang.org/x/image/draw"
	"image"
	"image/png"
	"path"
)

func zipEmojis(pd map[string]*types.ProviderDescriptor, mapScheme map[string]*typesScheme.Coordinates, pythonClient *gopy.Context) {
	for _, v := range pd {
		run, _ := pythonClient.Run(path.Join(consts.PythonLibPillow, "compress.py"), v.Emojis)
		_ = json.Unmarshal(run, &v.Emojis)
	}
	for _, v := range pd {
		v.EmojiCount = len(v.Emojis)
	}
	// Fix missing emojis
	for appleEmoji, appleContent := range pd["apple"].Emojis {
		for _, v := range pd {
			if _, ok := v.Emojis[appleEmoji]; !ok {
				v.Emojis[appleEmoji] = appleContent
			}
		}
	}
	for _, provider := range pd {
		var buf bytes.Buffer
		w := zip.NewWriter(&buf)
		for key, value := range provider.Emojis {
			s := mapScheme[key]
			fileName := fmt.Sprintf("%d_%d.png", s.X, s.Y)
			f, _ := w.Create(fileName)
			_, _ = f.Write(value)
		}
		_ = w.Close()
		provider.EmojiZip = buf.Bytes()
	}
	listPreview := []string{
		"😀",
		"😉",
		"😔",
		"😨",
	}
	emojiSize := 66
	for _, provider := range pd {
		previewImage := image.NewRGBA(image.Rect(0, 0, emojiSize*2, emojiSize*2))
		for i, emoji := range listPreview {
			emojiContent, _, _ := image.Decode(bytes.NewReader(provider.Emojis[emoji]))
			x := (i % 2) * emojiSize
			y := (i / 2) * emojiSize
			draw.Draw(previewImage, image.Rect(x, y, x+emojiSize, y+emojiSize), emojiContent, image.Point{}, draw.Src)
		}
		var buf bytes.Buffer
		_ = png.Encode(&buf, previewImage)
		provider.Preview = buf.Bytes()
	}
	for _, provider := range pd {
		packIdentifier := fmt.Sprintf("%d%d%d", provider.EmojiCount, len(provider.EmojiZip), provider.UnicodeVersion)
		byteSum := sha256.Sum256([]byte(packIdentifier))
		provider.MD5 = hex.EncodeToString(byteSum[:])
	}
}
