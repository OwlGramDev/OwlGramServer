package google_reviews

import (
	"OwlGramServer/consts"
	"context"
	"google.golang.org/api/androidpublisher/v3"
	"google.golang.org/api/option"
)

func Client() *Context {
	ctx := context.Background()
	pubService, err := androidpublisher.NewService(ctx, option.WithCredentialsFile(consts.GoogleApiJsonPath))
	if err != nil {
		return nil
	}
	return &Context{
		PlayServiceAPI_: pubService,
	}
}
