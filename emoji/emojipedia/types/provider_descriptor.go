package types

type ProviderDescriptor struct {
	Name           string `json:"name"`
	ID             string `json:"id"`
	Link           string
	Emojis         map[string][]byte
	EmojiZip       []byte
	Preview        []byte
	UnicodeVersion int
	EmojiCount     int
	MD5            string
}
