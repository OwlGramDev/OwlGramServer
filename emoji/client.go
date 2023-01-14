package emoji

import "OwlGramServer/gopy"

func Client(pythonClient *gopy.Context) *Context {
	return &Context{
		pythonClient: pythonClient,
	}
}
