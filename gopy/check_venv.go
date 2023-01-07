package gopy

func (c *Context) CheckVenv() bool {
	_, err := c.runWithVenv(nil)
	return err == nil
}
