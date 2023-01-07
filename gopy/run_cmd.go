package gopy

import (
	"bytes"
	"errors"
	"os/exec"
)

func runCmd(name string, args ...string) ([]byte, error) {
	var errMess bytes.Buffer
	cmd := exec.Command(name, args...)
	cmd.Stderr = &errMess
	res, err := cmd.Output()
	if err != nil {
		if errMess.Len() > 0 {
			return nil, errors.New(errMess.String())
		}
		return nil, err
	}
	return res, nil
}
