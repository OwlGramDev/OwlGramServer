package tg_checker

import "OwlGramServer/tg_checker/types"

type Context struct {
	StatusDC     []types.DCStatus
	LastRefresh  int64
	IsRefreshing bool
}
