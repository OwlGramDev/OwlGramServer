package gopy

func (c *Context) CheckVenv() bool {
	_, err := c.runWithVenv()
	return err == nil
}
