package gopy

import (
	"fmt"
	"strings"
)

func (c *Context) InstallPackages(packages ...string) error {
	_, err := c.runWithVenv(nil, fmt.Sprintf("pip install %s", strings.Join(packages, " ")))
	return err
}
