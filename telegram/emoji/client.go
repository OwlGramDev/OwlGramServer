package emoji

import (
	"OwlGramServer/consts"
	"context"
	"fmt"
	"github.com/gotd/td/telegram"
	"github.com/gotd/td/telegram/auth"
	"github.com/gotd/td/tgerr"
	"time"
)

func Client() *Context {
	session := &memorySession{}
	botClient := telegram.NewClient(consts.ApiID, consts.ApiHash, telegram.Options{
		NoUpdates:      true,
		SessionStorage: session,
	})
	waitClient := make(chan bool)
	var client *Context
	var phoneCodeHash string
	number := fmt.Sprintf("+%s%s", consts.PhoneCode, consts.PhoneUserbot)
	go func() {
		err := botClient.Run(context.Background(), func(ctx context.Context) error {
			status, err := botClient.Auth().Status(ctx)
			if err != nil {
				return err
			}
			if !status.Authorized {
			authCode:
				if len(phoneCodeHash) == 0 {
					result, err := botClient.Auth().SendCode(ctx, number, auth.SendCodeOptions{
						CurrentNumber: true,
					})
					nErr, _ := tgerr.As(err)
					if nErr != nil {
						if nErr.Type == "AUTH_RESTART" {
							goto authCode
						} else {
							_ = session.StoreSession(ctx, []byte{})
							return err
						}
					}
					phoneCodeHash = result.PhoneCodeHash
				}
			retryCode:
				fmt.Print("Enter the Sent Telegram code: ")
				var secretCode string
				_, _ = fmt.Scanln(&secretCode)
				in, err := botClient.Auth().SignIn(ctx, number, secretCode, phoneCodeHash)
				nErr, _ := tgerr.As(err)
				if nErr != nil {
					if nErr.Type == "PHONE_CODE_INVALID" {
						fmt.Println("Invalid code, try again")
						goto retryCode
					} else {
						return err
					}
				}
			login2FA:
				if in == nil {
					fmt.Print("Enter the 2FA password: ")
					var passwordInput string
					_, _ = fmt.Scanln(&passwordInput)
					in, _ = botClient.Auth().Password(ctx, passwordInput)
					if in == nil {
						fmt.Println("Invalid password, try again")
						goto login2FA
					}
				}
			}
			client = &Context{
				client:  botClient.API(),
				context: ctx,
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
