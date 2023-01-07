package utilities

import (
	"OwlGramServer/consts"
	"OwlGramServer/gopy"
	"fmt"
)

func SmartPythonBuild() *gopy.Context {
	pythonClient := gopy.Client("3.10", consts.VenvPath)
	if !pythonClient.CheckVenv() {
		fmt.Println(fmt.Sprintf("Building Python%s for venv...", pythonClient.GetVersion()))
		venvRes := pythonClient.BuildVenv()
		if venvRes != nil {
			panic(fmt.Sprintf("Error while building venv, error: %s", venvRes.Error()))
		}
		fmt.Println(fmt.Sprintf("Succefully builded for Python%s venv!", pythonClient.GetVersion()))
	}
	if !pythonClient.CheckInstallation(consts.Requirements...) {
		fmt.Println(fmt.Sprintf("Installing %d packages...", len(consts.Requirements)))
		err := pythonClient.InstallPackages(consts.Requirements...)
		if err != nil {
			panic(fmt.Sprintf("Error while installing packages, error: %s", err.Error()))
		}
		fmt.Println(fmt.Sprintf("Succefully installed %d packages!", len(consts.Requirements)))
	}
	return pythonClient
}
