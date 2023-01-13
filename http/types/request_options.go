package types

import "time"

type RequestOptions struct {
	Retries       int
	Timeout       time.Duration
	Method        string
	BearerToken   string
	Body          []byte
	Headers       map[string]string
	MultiPart     *MultiPartInfo
	NoInstantRead bool
}
