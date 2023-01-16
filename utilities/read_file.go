package utilities

import "io"

func ReadFile(file io.ReadCloser) ([]byte, error) {
	var buf []byte
	for {
		var b = make([]byte, 1024*4)
		n, fErr := io.ReadFull(file, b)
		buf = append(buf, b[:n]...)
		if fErr != nil {
			if fErr == io.EOF {
				break
			}
			if fErr != io.ErrUnexpectedEOF {
				return nil, fErr
			}
		}
	}
	return buf, nil
}
