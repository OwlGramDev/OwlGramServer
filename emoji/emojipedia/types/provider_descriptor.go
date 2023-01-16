package types

import "strings"

type ProviderDescriptor struct {
	Name           string            `json:"name"`
	ID             string            `json:"id"`
	Margin         int               `json:"margin"`
	Emojis         map[string][]byte `json:"emojis"`
	Link           string
	EmojiZip       []byte
	Preview        []byte
	UnicodeVersion int
	EmojiCount     int
	MD5            string
}

func (p *ProviderDescriptor) GetID() string {
	return strings.Split(p.ID, "/")[0]
}

func (p *ProviderDescriptor) HaveVariant() bool {
	return strings.Contains(p.ID, "/")
}

func (p *ProviderDescriptor) GetVariant() string {
	return strings.Split(p.ID, "/")[1]
}
