package gopy

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
)

func (c *Context) Run(path string, params any) ([]byte, error) {
	marshal, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}
	return c.runWithVenv(fmt.Sprintf("python3 %s %s", path, hex.EncodeToString(marshal)))
}
