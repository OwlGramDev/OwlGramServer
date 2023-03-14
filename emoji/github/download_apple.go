package github

import (
	"OwlGramServer/consts"
	typeScheme "OwlGramServer/emoji/scheme/types"
	"OwlGramServer/http"
	"OwlGramServer/utilities"
	"archive/zip"
	"bytes"
	"context"
	"github.com/bradleyfalzon/ghinstallation/v2"
	"github.com/google/go-github/v45/github"
	stdHttp "net/http"
	"path"
	"path/filepath"
	"strconv"
	"strings"
)

func DownloadApple(scheme *typeScheme.TgAScheme) (map[string][]byte, error) {
	emojiList := make(map[string][]byte)
	itr, err := ghinstallation.NewKeyFromFile(stdHttp.DefaultTransport, 163227, 22042191, consts.GithubPemFile)
	if err != nil {
		return nil, err
	}
	client := github.NewClient(&stdHttp.Client{Transport: itr})
	rc, _, _ := client.Repositories.GetArchiveLink(context.Background(), consts.GithubRepoOwner, consts.GithubRepo, github.Zipball, &github.RepositoryContentGetOptions{
		Ref: "develop",
	}, true)
	r := http.ExecuteRequest(rc.String())
	zipReader, err := zip.NewReader(bytes.NewReader(r.Read()), int64(len(r.Read())))
	if err != nil {
		return nil, err
	}
	unicodeAlias := make(map[string]string)
	for emoji := range scheme.Data {
		unicode := utilities.GetEmojiFromHex(utilities.GetHexFromEmoji(emoji, " ", false), " ", false)
		unicodeAlias[unicode] = emoji
	}
	for _, f := range zipReader.File {
		parent := strings.Split(f.Name, string(filepath.Separator))[0]
		if filepath.Join(parent, "TMessagesProj", "src", "main", "assets", "emoji") == filepath.Dir(f.Name) && !f.FileInfo().IsDir() {
			emojiFile, e := f.Open()
			if e != nil {
				return nil, e
			}
			buf, e := utilities.ReadFile(emojiFile)
			if e != nil {
				return nil, e
			}
			_ = emojiFile.Close()
			fileName := path.Base(f.Name)
			emojiData := strings.Split(fileName[:len(fileName)-len(path.Ext(fileName))], "_")
			x, _ := strconv.Atoi(emojiData[0])
			y, _ := strconv.Atoi(emojiData[1])
			emoji := scheme.GetEmoji(typeScheme.Coordinates{X: x, Y: y})
			if len(emoji) == 0 {
				continue
			}
			emojiList[emoji] = buf
		}
	}
	return emojiList, nil
}
