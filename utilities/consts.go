package utilities

const (
	NoZipFound = iota
	UnzipError
	NoBundleFound
	ReadingBundlesInfo
	BundlesHaveDifferentVersions
	BundleCorrupted
	BundlesCompiled
	BundlesSelectReleaseType
	SendingToStores
	FailedToSendToStores
	BundlesSentToStores
	SelectBaseBundle
	NeededChangelogs
	EditingChangelogs
	EditedChangelogs
	NeededImage
	NeededDesc
	NeededNotes
	NeededLink
	ConfirmChanges
	ChangesConfirmed
)

const (
	StatusPending     = "pending"
	StatusRunning     = "running"
	StatusUploading   = "uploading"
	StatusExtracting  = "extracting"
	StatusSuccess     = "success"
	StatusElaborating = "elaborating"
	StatusFailed      = "failed"
	StatusCanceled    = "canceled"
)
