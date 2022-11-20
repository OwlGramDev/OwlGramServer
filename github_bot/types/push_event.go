package types

type PushEvent struct {
	ID         string `json:"id"`
	Repository struct {
		FullName string `json:"full_name"`
		HTMLUrl  string `json:"html_url"`
	} `json:"repository"`
	Commits []struct {
		ID      string `json:"id"`
		Message string `json:"message"`
		URL     string `json:"url"`
	}
	Compare string `json:"compare"`
}
