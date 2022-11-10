package types

type UpdatesDescriptor struct {
	Localizations map[string]map[string]string `json:"localizations"`
	Updates       *UpdateList                  `json:"updates"`
}
