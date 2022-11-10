package types

type UpdateList struct {
	Stable *UpdateInfo `json:"stable"`
	Beta   *UpdateInfo `json:"beta"`
}
