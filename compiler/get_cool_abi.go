package compiler

func GetCoolAbi(abi string, typeStyle int) string {
	if typeStyle == 0 {
		switch abi {
		case "universal":
			return "Universal"
		case "armeabi-v7a":
			return "Armv7"
		case "arm64-v8a":
			return "Arm64"
		default:
			return abi
		}
	} else if typeStyle == 1 {
		switch abi {
		case "armeabi-v7a":
			return "arm-v7a"
		case "x64":
			return "x86_64"
		default:
			return abi
		}
	} else if typeStyle == 3 {
		switch abi {
		case "armeabi-v7a":
			return "armv7"
		case "arm64-v8a":
			return "arm64"
		default:
			return abi
		}
	} else {
		switch abi {
		case "armeabi-v7a":
			return "arm-v7a"
		default:
			return abi
		}
	}
}
