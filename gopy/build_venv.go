package gopy

import "fmt"

func (c *Context) BuildVenv() error {
	_, err := runCmd(nil, fmt.Sprintf("python%s", c.pythonVersion), "-m", "venv", c.venvPath)
	return err
}
