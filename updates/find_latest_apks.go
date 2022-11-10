package updates

import (
	"OwlGramServer/consts"
	"fmt"
	"os"
	"path"
	"strconv"
)

func (ctx *Context) findLatestApks(isBeta bool) string {
	pathFiles := path.Join(consts.ServerFilesFolder, "versions")
	f, err := os.Open(pathFiles)
	if err != nil {
		panic(fmt.Sprintf("Error opening folder %s", pathFiles))
		return ""
	}
	files, err := f.Readdir(0)
	if err != nil {
		panic("Error reading directory")
	}
	latestVersion := -1
	filePath := ""
	for _, v := range files {
		if v.IsDir() {
			var fileName string
			if isBeta {
				fileName = path.Join(pathFiles, v.Name(), "universal-beta.apk")
			} else {
				fileName = path.Join(pathFiles, v.Name(), "universal-stable.apk")
			}
			if _, err := os.Stat(fileName); os.IsNotExist(err) {
				continue
			}
			intVersion, err := strconv.Atoi(v.Name())
			if err != nil {
				continue
			}
			if intVersion > latestVersion {
				latestVersion = intVersion
				filePath = fileName
			}
		}
	}
	if latestVersion == -1 && isBeta {
		ctx.findLatestApks(false)
	}
	return filePath
}
