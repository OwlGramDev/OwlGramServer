package github

import (
	"OwlGramServer/consts"
	"OwlGramServer/telegram/emoji/github/types"
	"context"
	"github.com/google/go-github/v45/github"
	"sync"
)

func getPaths(ctx context.Context, client *github.Client, path string) []*types.FileDescriptor {
	_, contents, _, err := client.Repositories.GetContents(ctx, consts.GithubRepoOwnerZFont, consts.GithubRepoZFont, path, nil)
	if err != nil {
		return nil
	}
	var result []*types.FileDescriptor
	var waitSync sync.WaitGroup
	for _, content := range contents {
		waitSync.Add(1)
		go func(content *github.RepositoryContent) {
			if content.GetType() == "dir" {
				result = append(result, getPaths(ctx, client, content.GetPath())...)
			} else {
				commits, _, _ := client.Repositories.ListCommits(ctx, consts.GithubRepoOwnerZFont, consts.GithubRepoZFont, &github.CommitsListOptions{
					Path: content.GetPath(),
				})
				result = append(result, &types.FileDescriptor{
					Name:        content.GetName(),
					DownloadURL: content.GetDownloadURL(),
					LastUpdate:  commits[0].GetCommit().GetAuthor().GetDate(),
				})
			}
			waitSync.Done()
		}(content)
	}
	waitSync.Wait()
	return result
}
