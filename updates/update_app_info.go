package updates

import (
	"OwlGramServer/compiler"
	"OwlGramServer/updates/types"
)

func (ctx *Context) updateAppInfo(input *types.UpdatesDescriptor) {
	t, _ := compiler.GetApkInfo(ctx.findLatestApks(false))
	input.Updates.Stable.VersionName = t.VersionName
	input.Updates.Stable.VersionCode = t.VersionCode
	t, _ = compiler.GetApkInfo(ctx.findLatestApks(true))
	if len(t.Path) == 0 {
		input.Updates.Beta.VersionName = input.Updates.Stable.VersionName
		input.Updates.Beta.VersionCode = input.Updates.Stable.VersionCode
	} else {
		input.Updates.Beta.VersionName = t.VersionName
		input.Updates.Beta.VersionCode = t.VersionCode
	}
}
