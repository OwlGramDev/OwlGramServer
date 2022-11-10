package updates

func GetFile(isBeta bool, abi string) string {
	linkDownload := abi
	if isBeta {
		linkDownload += "-beta.apk"
	} else {
		linkDownload += "-stable.apk"
	}
	return linkDownload
}
