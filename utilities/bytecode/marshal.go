package bytecode

import (
	"bytes"
	"compress/gzip"
	"encoding/gob"
)

func Marshal(input any) ([]byte, error) {
	// Marshal Data
	buf := bytes.Buffer{}
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(input)
	if err != nil {
		return nil, err
	}

	// Compress Data
	zipBuf := bytes.Buffer{}
	zipped := gzip.NewWriter(&zipBuf)
	_, err = zipped.Write(buf.Bytes())
	if err != nil {
		return nil, err
	}
	err = zipped.Close()
	if err != nil {
		return nil, err
	}
	return zipBuf.Bytes(), nil
}
