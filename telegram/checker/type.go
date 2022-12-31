package checker

import "OwlGramServer/telegram/checker/types"

type Context struct {
	StatusDC     []types.DCStatus
	LastRefresh  int64
	IsRefreshing bool
}
