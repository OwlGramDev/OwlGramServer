package types

type RequestOptions struct {
	Retries     int
	Method      string
	BearerToken string
	Body        []byte
	Headers     map[string]string
	MultiPart   *MultiPartInfo
}
