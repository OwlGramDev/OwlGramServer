package types

import (
	"io"
)

type HTTPResult struct {
	Body       io.ReadCloser
	Error      error
	StatusCode int
	cacheRead  []byte
}

func (r *HTTPResult) SetFallback(body []byte) {
	r.cacheRead = body
	r.Body = nil
}

func (r *HTTPResult) Read() []byte {
	if r.cacheRead != nil {
		return r.cacheRead
	}
	if r.Body == nil {
		return nil
	}
	var buf []byte
	for {
		var b = make([]byte, 1024*4)
		n, err := io.ReadFull(r.Body, b)
		buf = append(buf, b[:n]...)
		if err != nil {
			if err == io.EOF {
				break
			}
			if err != io.ErrUnexpectedEOF {
				return nil
			}
		}
	}
	defer func() {
		r.close()
		r.Body = nil
	}()
	r.cacheRead = buf
	return r.cacheRead
}

func (r *HTTPResult) ReadString() string {
	return string(r.Read())
}

func (r *HTTPResult) close() {
	_ = r.Body.Close()
}
