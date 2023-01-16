package bytecode

import (
	"bytes"
	"compress/gzip"
	"encoding/gob"
	"io"
)

func Unmarshal(data []byte, v any) error {
	// Uncompress Data
	rdr, _ := gzip.NewReader(bytes.NewReader(data))
	data, err := io.ReadAll(rdr)
	if err != nil {
		return err
	}
	err = rdr.Close()
	if err != nil {
		return err
	}

	// Unmarshal Data
	dec := gob.NewDecoder(bytes.NewReader(data))
	err = dec.Decode(v)
	if err != nil {
		return err
	}
	return nil
}
