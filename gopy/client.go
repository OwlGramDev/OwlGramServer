package gopy

func Client(pythonVersion, venvPath string) *Context {
	return &Context{
		venvPath:      venvPath,
		pythonVersion: pythonVersion,
	}
}
