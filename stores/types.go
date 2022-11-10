package stores

import (
	"OwlGramServer/stores/types"
	"OwlGramServer/updates"
)

type Context struct {
	StoreList    []*types.StoreInfo
	updateClient *updates.Context
}
