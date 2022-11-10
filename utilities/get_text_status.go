package utilities

func GetTextStatus(status string) string {
	switch status {
	case StatusPending:
		return "In attesa..."
	case StatusRunning:
		return "In corso..."
	case StatusSuccess:
		return "Completato"
	case StatusFailed:
		return "Fallito"
	case StatusCanceled:
		return "Annullato"
	case StatusExtracting:
		return "Estrazione in corso..."
	case StatusUploading:
		return "Upload in corso..."
	case StatusElaborating:
		return "Elaborazione in corso..."
	default:
		return "Sconosciuto"
	}
}
