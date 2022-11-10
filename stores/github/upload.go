package github

import (
	"OwlGramServer/compiler"
	bundles "OwlGramServer/compiler/types"
	"OwlGramServer/consts"
	"OwlGramServer/stores/github/types"
	"context"
	"fmt"
	"github.com/bradleyfalzon/ghinstallation/v2"
	"github.com/google/go-github/v43/github"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
)

func Upload(apkList []bundles.PackageInfo, releaseType string, descriptionUpdate string, listener func(float64)) {
	var listApksTmp []bundles.PackageInfo
	for _, apk := range apkList {
		if apk.IsApk {
			listApksTmp = append(listApksTmp, apk)
		}
	}
	apkList = listApksTmp
	isReUpload := strings.Split(releaseType, "-")[0] == "re"
	if isReUpload {
		releaseType = strings.Split(releaseType, "-")[1]
	}
	ctx := context.Background()
	itr, err := ghinstallation.NewKeyFromFile(http.DefaultTransport, 163227, 22042191, consts.GithubPemFile)
	if err != nil {
		listener(-1)
		return
	}
	client := github.NewClient(&http.Client{Transport: itr})
	releases, _, err := client.Repositories.ListReleases(ctx, consts.GithubRepoOwner, consts.GithubRepo, nil)
	if err != nil {
		listener(-1)
		return
	}
	var release *github.RepositoryRelease
	for _, r := range releases {
		if r.GetTagName() == fmt.Sprintf("#%s", releaseType) {
			release = r
			break
		}
	}
	if release != nil {
		_, _ = client.Repositories.DeleteRelease(ctx, consts.GithubRepoOwner, consts.GithubRepo, release.GetID())
	}
	re := regexp.MustCompile(`<[^>]*>?`)
	descriptionUpdate = re.ReplaceAllString(descriptionUpdate, "")
	descriptionUpdate = strings.Replace(descriptionUpdate, strings.Split(descriptionUpdate, "\n\n")[0]+"\n\n", "", -1)
	releaseTmp := types.Release{
		TagName:    fmt.Sprintf("#%s", releaseType),
		Name:       fmt.Sprintf("💡 Owlgram %s (%d)", apkList[0].VersionName, apkList[0].VersionCode),
		Body:       descriptionUpdate,
		PreRelease: releaseType == "beta",
	}
	release, res, err := client.Repositories.CreateRelease(ctx, consts.GithubRepoOwner, consts.GithubRepo, &github.RepositoryRelease{
		TagName:    &releaseTmp.TagName,
		Name:       &releaseTmp.Name,
		Body:       &releaseTmp.Body,
		Prerelease: &releaseTmp.PreRelease,
	})
	if err != nil || release.ID == nil || res.StatusCode != 201 {
		listener(-1)
		return
	}
	for i := 0; i < len(apkList); i++ {
		file, err := os.Open(apkList[i].Path)
		if err != nil {
			log.Println(err)
			listener(-1)
			return
		}
		apkName := fmt.Sprintf("OwlGram %s %s", apkList[i].VersionName, compiler.GetCoolAbi(apkList[i].AbiName, 1))
		_, res, err = client.Repositories.UploadReleaseAsset(ctx, consts.GithubRepoOwner, consts.GithubRepo, *release.ID, &github.UploadOptions{
			Name:  fmt.Sprintf("%s.apk", apkName),
			Label: apkName,
		}, file)
		if err != nil || res.StatusCode != 201 {
			listener(-1)
			return
		}
		if isReUpload {
			listener(50 + ((float64(i+1) / float64(len(apkList)) * 100) / 2))
		} else if releaseType == "stable" {
			listener((float64(i+1) / float64(len(apkList)) * 100) / 2)
		} else {
			listener(float64(i+1) / float64(len(apkList)) * 100)
		}
	}
	if releaseType == "stable" {
		Upload(apkList, "re-beta", descriptionUpdate, listener)
	} else {
		listener(-2)
	}
}
