package gopy

import (
	"fmt"
	"path"
	"strings"
)

func (c *Context) runWithVenv(args ...string) ([]byte, error) {
	args = append([]string{fmt.Sprintf("source %s", path.Join(c.venvPath, "/bin/activate"))}, args...)
	return runCmd(
		"bash",
		"-c",
		strings.Join(args, " && "),
	)
}
