package utilities

import "fmt"

func FormatSize(size uint64, removeZero bool) string {
	if size < 1024 {
		return fmt.Sprintf("%d B", size)
	} else if size < 1024*1024 {
		value := float64(size) / 1024.0
		if removeZero && (value-float64(uint64(value))*10 == 0) {
			return fmt.Sprintf("%d KB", uint64(value))
		} else {
			return fmt.Sprintf("%.1f KB", value)
		}
	} else if size < 1024*1024*1024 {
		value := float64(size) / 1024.0 / 1024.0
		if removeZero && (value-float64(uint64(value))*10 == 0) {
			return fmt.Sprintf("%d MB", uint64(value))
		} else {
			return fmt.Sprintf("%.1f MB", value)
		}
	} else {
		value := float64(size) / 1024.0 / 1024.0 / 1024.0
		if removeZero && (value-float64(uint64(value))*10 == 0) {
			return fmt.Sprintf("%d GB", uint64(value))
		} else {
			return fmt.Sprintf("%.2f GB", value)
		}
	}
}
