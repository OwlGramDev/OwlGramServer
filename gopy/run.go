package gopy

import (
	"encoding/json"
	"fmt"
)

func (c *Context) Run(path string, params any) ([]byte, error) {
	marshal, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}
	return c.runWithVenv(marshal, fmt.Sprintf("python %s", path))
}
