package compiler

func GetAbiName(abiCode int) string {
	switch abiCode {
	case 1, 3:
		return "armeabi-v7a"
	case 2, 4:
		return "x64"
	case 5, 7:
		return "arm64-v8a"
	case 6, 8:
		return "x86"
	default:
		return "universal"
	}
}
