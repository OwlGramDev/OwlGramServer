package compiler

import (
	"os/exec"
)

func UnzipFiles(zipFile, dest string) error {
	cmd := exec.Command("unzip", "-o", zipFile, "-d", dest)
	return cmd.Run()
}
