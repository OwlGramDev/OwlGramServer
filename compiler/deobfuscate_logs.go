package compiler

import (
	"OwlGramServer/consts"
	"bytes"
	"errors"
	"os"
	"os/exec"
	"path"
	"strconv"
)

func DeObfuscateLogs(logData string, appVersion int) ([]byte, error) {
	pathMapping := path.Join(consts.ProGuardFolder, strconv.Itoa(appVersion)+"-mapping.txt")
	pathInput := path.Join(consts.CacheFolder, "logs.txt")
	_ = os.WriteFile(path.Join(consts.CacheFolder, "logs.txt"), []byte(logData), 0775)
	cmd := exec.Command("java", "-jar", consts.RetraceToolPath, pathMapping, pathInput)
	var stdOut, stdErr bytes.Buffer
	cmd.Stdout = &stdOut
	cmd.Stderr = &stdErr
	if err := cmd.Run(); err != nil {
		return nil, errors.New(string(stdErr.Bytes()))
	}
	return stdOut.Bytes(), nil
}
