package types

import (
	"bytes"
	"io"
)

type HTTPResult struct {
	Body      io.ReadCloser
	Error     error
	cacheRead []byte
}

func (r *HTTPResult) Read() []byte {
	if r.cacheRead != nil {
		return r.cacheRead
	}
	defer r.Close()
	var buf bytes.Buffer
	_, err := io.Copy(&buf, r.Body)
	if err != nil {
		return nil
	}
	r.cacheRead = buf.Bytes()
	return r.cacheRead
}

func (r *HTTPResult) ReadString() string {
	return string(r.Read())
}

func (r *HTTPResult) Close() {
	_ = r.Body.Close()
}
