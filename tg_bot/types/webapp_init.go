package types

type WebAppInit struct {
	QueryId string `json:"query_id"`
	User    struct {
		ID        int64  `json:"id"`
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Username  string `json:"username"`
		Language  string `json:"language_code"`
		IsPremium bool   `json:"is_premium,omitempty"`
	} `json:"user"`
	AuthDate string `json:"auth_date"`
	Hash     string `json:"hash"`
}
