package compiler

import (
	"path"
	"strings"
)

func orderBundles(bundles []string) []string {
	listReturn := make([]string, 5)
	for bundle := range bundles {
		if path.Ext(bundles[bundle]) == ".apk" {
			if strings.Contains(bundles[bundle], "universal") {
				listReturn[0] = bundles[bundle]
			} else if strings.Contains(bundles[bundle], "arm64-v8a") {
				listReturn[1] = bundles[bundle]
			} else if strings.Contains(bundles[bundle], "armeabi-v7a") {
				listReturn[2] = bundles[bundle]
			} else if strings.Contains(bundles[bundle], "x86_64") {
				listReturn[3] = bundles[bundle]
			} else if strings.Contains(bundles[bundle], "x86") {
				listReturn[4] = bundles[bundle]
			}
		} else {
			listReturn = append(listReturn, bundles[bundle])
		}
	}
	var newList []string
	for i := range listReturn {
		if listReturn[i] != "" {
			newList = append(newList, listReturn[i])
		}
	}
	return newList
}
