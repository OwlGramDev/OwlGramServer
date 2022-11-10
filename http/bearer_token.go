package http

import "OwlGramServer/http/types"

type bearerTokenOption string

func (ct bearerTokenOption) Apply(o *types.RequestOptions) {
	o.BearerToken = string(ct)
}

func BearerToken(method string) RequestOption {
	return bearerTokenOption(method)
}
