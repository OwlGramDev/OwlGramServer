package emoji

import (
	"OwlGramServer/telegram/emoji/scheme/types"
	"OwlGramServer/utilities"
	"archive/zip"
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/gotd/td/tg"
	"image"
	"image/png"
	"io"
	"math"
	"strconv"
	"strings"
)

type Pack struct {
	DisplayName string
	Version     int
	Position    int
	Identifier  string
	Preview     []byte
	Emojies     []byte
	MD5         string
	previewDoc  *tg.MessageMediaDocument
	emojiesDoc  *tg.MessageMediaDocument
}

func newEmojiPack(message *tg.Message) *Pack {
	var version int
	var displayName string
	var position int
	for _, entity := range message.Entities {
		text := message.Message[entity.GetOffset() : entity.GetOffset()+entity.GetLength()]
		if utilities.InstanceOf(entity, &tg.MessageEntityHashtag{}) {
			hashtag := text[1:]
			if strings.HasPrefix(hashtag, "v") {
				version, _ = strconv.Atoi(hashtag[1:])
			} else if strings.HasPrefix(hashtag, "p") {
				position, _ = strconv.Atoi(hashtag[1:])
			}
		} else if utilities.InstanceOf(entity, &tg.MessageEntityCode{}) {
			displayName = text
		}
	}
	media := message.Media.(*tg.MessageMediaDocument)
	return &Pack{
		DisplayName: displayName,
		Version:     version,
		Position:    position,
		Identifier:  getIdentifier(media),
		emojiesDoc:  media,
	}
}

func (e *Pack) SetPreview(preview *tg.MessageMediaDocument) {
	e.previewDoc = preview
}

func (e *Pack) Download(c *Context) {
	for i := 0; i < 3; i++ {
		e.Preview, _ = c.downloadBytes(e.previewDoc)
		if e.Preview != nil {
			break
		}
	}
	for i := 0; i < 3; i++ {
		e.Emojies, _ = c.downloadBytes(e.emojiesDoc)
		if e.Emojies != nil {
			break
		}
	}
}

func (e *Pack) BuildPack(schemeLayout *types.NewScheme) {
	e.zipEmojis(e.unpackAsset(), e.mapSprites(schemeLayout))
}

func (e *Pack) mapSprites(schemeLayout *types.NewScheme) map[string]*sprite {
	emojiSprites := make(map[string]*sprite)
	emojiOriginalSize := int(30 * schemeLayout.Scale)
	for sectionIndex := range schemeLayout.Data {
		count2 := int(math.Ceil(float64(len(schemeLayout.Data[sectionIndex])) / float64(schemeLayout.SplitCount)))
		for emojiIndex := range schemeLayout.Data[sectionIndex] {
			emoji := schemeLayout.Data[sectionIndex][emojiIndex]
			if emoji == nil {
				continue
			}
			page := emojiIndex / count2
			position := emojiIndex - page*count2
			row := position % schemeLayout.Columns[sectionIndex][page]
			col := position / schemeLayout.Columns[sectionIndex][page]
			margin := int(float64(schemeLayout.Margins[sectionIndex][page]) * schemeLayout.Scale)
			marginLeft := margin * row
			marginTop := margin * col

			left := row*emojiOriginalSize + marginLeft
			top := col*emojiOriginalSize + marginTop
			emojiSprites[emoji.Emoji] = &sprite{
				SectionIndex: sectionIndex,
				Page:         page,
				Rect:         image.Rect(left, top, left+emojiOriginalSize, top+emojiOriginalSize),
				Coordinates:  emoji.Coordinates,
			}
		}
	}
	return emojiSprites
}

func (e *Pack) zipEmojis(emojiAsset map[string][]byte, sprites map[string]*sprite) {
	assets := make(map[int]map[int]image.Image)
	for key, value := range emojiAsset {
		data := strings.Split(key, "_")
		sectionIndex, _ := strconv.Atoi(data[0])
		page, _ := strconv.Atoi(data[1])
		if assets[sectionIndex] == nil {
			assets[sectionIndex] = make(map[int]image.Image)
		}
		assets[sectionIndex][page], _ = png.Decode(bytes.NewReader(value))
	}
	var buf bytes.Buffer
	w := zip.NewWriter(&buf)
	for _, s := range sprites {
		img := assets[s.SectionIndex][s.Page]
		emojiImage := img.(*image.Paletted).SubImage(s.Rect)
		var buffer bytes.Buffer
		_ = png.Encode(&buffer, emojiImage)
		fileName := fmt.Sprintf("%d_%d.png", s.Coordinates.X, s.Coordinates.Y)
		f, _ := w.Create(fileName)
		_, _ = f.Write(buffer.Bytes())
	}
	_ = w.Close()
	e.Emojies = buf.Bytes()
	packIdentifier := fmt.Sprintf("%d%d%d%d", len(sprites), len(assets), len(e.Emojies), e.Version)
	byteSum := sha256.Sum256([]byte(packIdentifier))
	e.MD5 = hex.EncodeToString(byteSum[:])
}

func (e *Pack) unpackAsset() map[string][]byte {
	emojiAsset := make(map[string][]byte)
	zipReader := bytes.NewReader(e.Emojies)
	rdr, _ := zip.NewReader(zipReader, int64(len(e.Emojies)))
	for _, file := range rdr.File {
		zipFileReader, _ := file.Open()
		var buf bytes.Buffer
		_, _ = io.Copy(&buf, zipFileReader)
		_ = zipFileReader.Close()
		fileName := file.Name
		if strings.Contains(fileName, ".") {
			fileName = fileName[:strings.LastIndex(fileName, ".")]
		}
		emojiAsset[fileName] = buf.Bytes()
	}
	return emojiAsset
}
