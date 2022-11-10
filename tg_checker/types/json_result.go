package types

type JSONResult struct {
	Status        []DCStatus `json:"status"`
	LastRefresh   int64      `json:"last_refresh"`
	RefreshInTime int8       `json:"refresh_in_time"`
	IsRefreshing  bool       `json:"is_refreshing"`
}
