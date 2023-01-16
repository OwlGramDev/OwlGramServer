package emoji

func LegacyID(idFull, idShort string) string {
	switch idFull {
	case "microsoft-teams":
		return "fluent"
	case "google/android-7.1":
		return "blob"
	default:
		return idShort
	}
}
