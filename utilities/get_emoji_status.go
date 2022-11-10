package utilities

func GetEmojiStatus(status string) string {
	switch status {
	case StatusPending:
		return "🕑"
	case StatusRunning:
		return "🛠"
	case StatusSuccess:
		return "✔️"
	case StatusFailed:
		return "🚫"
	case StatusCanceled:
		return "✖️"
	case StatusExtracting:
		return "📦"
	case StatusUploading:
		return "📤"
	case StatusElaborating:
		return "📝"
	default:
		return "❓"
	}
}
