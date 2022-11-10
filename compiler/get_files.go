package compiler

import (
	"os"
	"path"
)

func getFiles(folder string) []string {
	var listFiles []string
	f, _ := os.Open(folder)
	files, _ := f.Readdir(0)
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		switch path.Ext(file.Name()) {
		case ".apk", ".aab":
			listFiles = append(listFiles, folder+"/"+file.Name())
		}
	}
	return listFiles
}
