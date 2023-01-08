package github

import (
	"OwlGramServer/telegram/emoji/github/types"
	typeScheme "OwlGramServer/telegram/emoji/scheme/types"
	"archive/zip"
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"image"
	"image/draw"
	"image/png"
)

func buildSprites(appleEmoji map[int]map[int][]byte, schemeLayout *typeScheme.NewScheme, data []*types.FileDescriptor) ([]*types.PackTMP, error) {
	var packs []*types.PackTMP
	for _, file := range data {
		e := &types.PackTMP{
			DisplayName: file.Name,
			Identifier:  file.ID,
		}
		var buf bytes.Buffer
		w := zip.NewWriter(&buf)
		previewsEmojis := map[string]image.Image{
			"😀": nil,
			"😉": nil,
			"😔": nil,
			"😨": nil,
		}
		var totalItem int
		for page, pageData := range schemeLayout.Data {
			for section, s := range pageData {
				if s == nil {
					continue
				}
				fileName := fmt.Sprintf("%d_%d.png", s.Coordinates.X, s.Coordinates.Y)
				sprite := file.EmojiSprites[page][section]
				if sprite != nil {
					totalItem++
				} else {
					sprite = appleEmoji[s.Coordinates.X][s.Coordinates.Y]
					if sprite == nil {
						continue
					}
				}
				f, _ := w.Create(fileName)
				_, _ = f.Write(sprite)
				if _, ok := previewsEmojis[s.Emoji]; ok {
					if sprite == nil {
						return nil, fmt.Errorf("sprite is nil")
					}
					previewsEmojis[s.Emoji], _, _ = image.Decode(bytes.NewReader(sprite))
				}
			}
		}
		_ = w.Close()
		e.Emojies = buf.Bytes()

		previewImage := image.NewRGBA(image.Rect(0, 0, 132, 132))
		curr := 0
		emojiSize := 66
		for _, preview := range previewsEmojis {
			if preview == nil {
				return nil, fmt.Errorf("preview is nil")
			}
			x := (curr % 2) * emojiSize
			y := (curr / 2) * emojiSize
			draw.Draw(previewImage, image.Rect(x, y, x+emojiSize, y+emojiSize), preview, image.Point{}, draw.Src)
			curr++
		}
		var buffer bytes.Buffer
		_ = png.Encode(&buffer, previewImage)
		e.Preview = buffer.Bytes()
		md5Build := file.Content
		packIdentifier := fmt.Sprintf("%d%d", len(e.Emojies), totalItem)
		md5Build = append(md5Build, packIdentifier...)
		byteSum := sha256.Sum256(md5Build)
		e.MD5 = hex.EncodeToString(byteSum[:])
		packs = append(packs, e)
	}
	return packs, nil
}
