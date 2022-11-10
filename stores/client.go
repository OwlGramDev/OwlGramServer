package stores

import "OwlGramServer/updates"

func Client(updateClient *updates.Context) *Context {
	return &Context{
		updateClient: updateClient,
	}
}
