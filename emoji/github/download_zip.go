package github

import (
	"OwlGramServer/consts"
	"OwlGramServer/emoji/github/types"
	typeScheme "OwlGramServer/emoji/scheme/types"
	"OwlGramServer/http"
	"OwlGramServer/utilities"
	"archive/zip"
	"bytes"
	"context"
	"encoding/json"
	"github.com/bradleyfalzon/ghinstallation/v2"
	"github.com/google/go-github/v45/github"
	stdHttp "net/http"
	"path"
	"strings"
)

func DownloadZip(scheme *typeScheme.TgAScheme) (map[string][]byte, error) {
	itr, err := ghinstallation.NewKeyFromFile(stdHttp.DefaultTransport, 163227, 22042191, consts.GithubPemFile)
	if err != nil {
		return nil, err
	}
	client := github.NewClient(&stdHttp.Client{Transport: itr})
	rc, _, _ := client.Repositories.GetArchiveLink(context.Background(), consts.FluentGithubRepoOwner, consts.FluentGithubRepo, github.Zipball, &github.RepositoryContentGetOptions{
		Ref: "main",
	}, true)
	r := http.ExecuteRequest(rc.String())
	zipReader, _ := zip.NewReader(bytes.NewReader(r.Read()), int64(len(r.Read())))
	unicodeAlias := make(map[string]string)
	for emoji := range scheme.Data {
		unicode := utilities.GetEmojiFromHex(utilities.GetHexFromEmoji(emoji, " ", false), " ", false)
		unicodeAlias[unicode] = emoji
	}
	emojiTones := []string{
		"Default",
		"Light",
		"Medium-Light",
		"Medium",
		"Medium-Dark",
		"Dark",
	}
	emojiList := make(map[string]string)
	for _, f := range zipReader.File {
		if strings.HasSuffix(f.Name, "metadata.json") {
			metadata, e := f.Open()
			if e != nil {
				return nil, e
			}
			var metadataMap types.EmojiMetaData
			buf, fErr := utilities.ReadFile(metadata)
			if fErr != nil {
				return nil, fErr
			}
			fErr = json.Unmarshal(buf, &metadataMap)
			if fErr != nil {
				return nil, fErr
			}
			_ = metadata.Close()
			emojiStandard := utilities.GetEmojiFromHex(metadataMap.Unicode, " ", false)
			if _, ok := unicodeAlias[emojiStandard]; !ok {
				continue
			}
			filePath := path.Join(f.Name, "..")
			if len(metadataMap.UnicodeSkinTones) > 0 {
				for i, skinTone := range metadataMap.UnicodeSkinTones {
					emojiWithSkinTone := utilities.GetEmojiFromHex(skinTone, " ", false)
					if _, ok := unicodeAlias[emojiWithSkinTone]; !ok {
						continue
					}
					emojiList[emojiWithSkinTone] = path.Join(filePath, emojiTones[i], "3D")
				}
			} else {
				emojiList[unicodeAlias[emojiStandard]] = path.Join(filePath, "3D")
			}
		}
	}
	emojiBytes := make(map[string][]byte)
	for _, f := range zipReader.File {
		for emoji, filePath := range emojiList {
			if strings.HasPrefix(f.Name, filePath) && strings.HasSuffix(f.Name, ".png") {
				file, e := f.Open()
				if e != nil {
					return nil, e
				}
				emojiBytes[emoji], _ = utilities.ReadFile(file)
				_ = file.Close()
			}
		}
	}
	return emojiBytes, err
}
