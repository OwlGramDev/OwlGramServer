package compiler

import (
	"OwlGramServer/consts"
	"archive/zip"
	"io"
	"os"
	"path"
	"strconv"
)

func ExtractMapping(bundlePath string, versionCode int) {
	r, _ := zip.OpenReader(bundlePath)
	defer func(r *zip.ReadCloser) {
		_ = r.Close()
	}(r)

	for _, f := range r.File {
		if f.Name != path.Join("BUNDLE-METADATA", "com.android.tools.build.obfuscation", "proguard.map") {
			continue
		}
		rc, _ := f.Open()
		if _, err := os.Stat(consts.ProGuardFolder); os.IsNotExist(err) {
			_ = os.Mkdir(consts.ProGuardFolder, 0775)
		}
		w, _ := os.Create(path.Join(consts.ProGuardFolder, strconv.Itoa(versionCode)+"-mapping.txt"))
		_, _ = io.Copy(w, rc)
		_ = rc.Close()
		break
	}
}
