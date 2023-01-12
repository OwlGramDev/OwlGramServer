package github

import (
	"OwlGramServer/consts"
	customHttp "OwlGramServer/http"
	"OwlGramServer/telegram/emoji/github/types"
	"context"
	"github.com/bradleyfalzon/ghinstallation/v2"
	"github.com/google/go-github/v45/github"
	"golang.org/x/exp/slices"
	"net/http"
	"strings"
)

func GetLatestEmojis() []*types.FileDescriptor {
	ctx := context.Background()
	itr, err := ghinstallation.NewKeyFromFile(http.DefaultTransport, 163227, 22042191, consts.GithubPemFile)
	if err != nil {
		return nil
	}
	client := github.NewClient(&http.Client{Transport: itr})
	files := getPaths(ctx, client, "Fonts/Emoji")
	slices.SortFunc(files, func(i, j *types.FileDescriptor) bool {
		return i.LastUpdate.After(j.LastUpdate)
	})
	var emojiFonts []*types.FileDescriptor
	for _, allowedFont := range consts.WhiteListFont {
		for _, file := range files {
			if strings.HasPrefix(strings.ToLower(file.Name), strings.ToLower(allowedFont.MatchName)) {
				file.Name = allowedFont.Name
				file.ID = strings.ToLower(strings.Split(allowedFont.Name, " ")[0])
				emojiFonts = append(emojiFonts, file)
				break
			}
		}
	}
	for _, file := range emojiFonts {
		file.Content = customHttp.ExecuteRequest(file.DownloadURL).Read()
	}
	return emojiFonts
}
