package emoji

import (
	"context"
	"errors"
	"github.com/gotd/td/telegram/auth"
	"github.com/gotd/td/tg"
)

type noSignUp struct{}

func (c noSignUp) SignUp(_ context.Context) (auth.UserInfo, error) {
	return auth.UserInfo{}, errors.New("not implemented")
}

func (c noSignUp) AcceptTermsOfService(_ context.Context, tos tg.HelpTermsOfService) error {
	return &auth.SignUpRequired{TermsOfService: tos}
}
