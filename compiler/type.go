package compiler

import "OwlGramServer/compiler/types"

type Context struct {
	IsRunning bool
	MessageID int64
	Bundles   []types.PackageInfo
}
