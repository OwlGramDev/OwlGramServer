package emoji

import (
	"OwlGramServer/consts"
	"OwlGramServer/gopy"
	"context"
	"fmt"
	"github.com/gotd/td/telegram"
	"github.com/gotd/td/telegram/auth"
	"time"
)

func Client(pythonClient *gopy.Context) *Context {
	session := &memorySession{}
	botClient := telegram.NewClient(consts.ApiID, consts.ApiHash, telegram.Options{
		NoUpdates:      true,
		SessionStorage: session,
	})
	waitClient := make(chan bool)
	var client *Context
	flow := auth.NewFlow(
		termAuth{phone: fmt.Sprintf("+%s%s", consts.PhoneCode, consts.PhoneUserbot)},
		auth.SendCodeOptions{},
	)
	go func() {
		err := botClient.Run(context.Background(), func(ctx context.Context) error {
			if err := botClient.Auth().IfNecessary(ctx, flow); err != nil {
				return err
			}
			client = &Context{
				client:       botClient.API(),
				context:      ctx,
				pythonClient: pythonClient,
			}
			waitClient <- true
			for {
				time.Sleep(time.Second * 1)
			}
		})
		if err != nil {
			panic(err)
		}
	}()
	<-waitClient
	return client
}
