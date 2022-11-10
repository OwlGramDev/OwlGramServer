package http

import "OwlGramServer/http/types"

type RequestOption interface {
	Apply(o *types.RequestOptions)
}

type MultiPartOption interface {
	Apply(o *types.MultiPartInfo)
}
