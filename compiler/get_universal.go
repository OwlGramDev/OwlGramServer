package compiler

import "OwlGramServer/compiler/types"

func GetUniversal(listBundles []types.PackageInfo) *types.PackageInfo {
	for _, bundle := range listBundles {
		if bundle.AbiName == "universal" && bundle.IsApk {
			return &bundle
		}
	}
	return nil
}
