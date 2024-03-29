package github

import (
	"OwlGramServer/consts"
	"OwlGramServer/handlers"
	"OwlGramServer/telegram/github/types"
	"encoding/json"
	"fmt"
	"github.com/Squirrel-Network/gobotapi"
	"github.com/Squirrel-Network/gobotapi/logger"
	"github.com/Squirrel-Network/gobotapi/methods"
	"github.com/valyala/fasthttp"
	"golang.org/x/exp/slices"
	"sync"
)

func SendPushEvent(ctx *fasthttp.RequestCtx) {
	token := ctx.Request.URI().QueryArgs().Peek("token")
	body := ctx.PostBody()
	if token == nil || string(token) != consts.GithubToken || body == nil {
		handlers.Forbidden(ctx)
		return
	}
	var event types.PushEvent
	err := json.Unmarshal(body, &event)
	if err != nil {
		handlers.Forbidden(ctx)
		return
	}
	branchName := event.Ref[11:]
	if len(event.Commits) == 0 {
		handlers.Forbidden(ctx)
		return
	}
	message := fmt.Sprintf("📋 New Update in <a href='%s/tree/%s'>%s/%s</a>\n\n", event.Repository.HTMLUrl, branchName, event.Repository.FullName, branchName)
	for i, commit := range event.Commits {
		diff := len(event.Commits) - (i + 1)
		if i >= 15 && len(event.Commits) > 15 && diff > 0 {
			moreCompare := fmt.Sprintf("%s/compare/%s...%s", event.Repository.HTMLUrl, commit.ID[:12], event.Commits[len(event.Commits)-1].ID[:12])
			message += fmt.Sprintf("➕ And %d more <a href='%s'>commits...</a>\n", diff, moreCompare)
			break
		}
		message += fmt.Sprintf("‣ %s (<a href='%s'>%s</a>)\n", commit.Message, commit.URL, commit.ID[:7])
	}
	message += fmt.Sprintf("\n<a href='%s'>🔨 %d new commits</a>", event.Compare, len(event.Commits))
	client := gobotapi.NewClient(consts.GithubBotToken)
	client.NoUpdates = true
	client.LoggingLevel = logger.Error
	client.SleepThreshold = 60
	_ = client.Start()
	wait := sync.WaitGroup{}
	for _, group := range consts.GithubGroups {
		if slices.Contains(group.AllowedBranches, branchName) {
			wait.Add(1)
			go func(group types.Group) {
				_, _ = client.Invoke(&methods.SendMessage{
					ChatID:           group.ID,
					Text:             message,
					ReplyToMessageID: group.ForumID,
					ParseMode:        "HTML",
				})
				wait.Done()
			}(group)
		}
	}
	wait.Wait()
	client.Stop()
}
