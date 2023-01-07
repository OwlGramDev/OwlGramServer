package gopy

import (
	"fmt"
	"strings"
)

func (c *Context) CheckInstallation(packages ...string) bool {
	res, err := c.runWithVenv(nil, fmt.Sprintf("pip freeze | grep -i -E '%s'", strings.Join(packages, "|")))
	if err != nil {
		return false
	}
	packagesList := strings.Split(string(res), "\n")
	packagesList = packagesList[:len(packagesList)-1]
	return len(packagesList) == len(packages)
}
