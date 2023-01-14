package gopy

import "fmt"

func (c *Context) RunRaw(path string, stdin []byte) ([]byte, error) {
	return c.runWithVenv(stdin, fmt.Sprintf("python %s", path))
}
