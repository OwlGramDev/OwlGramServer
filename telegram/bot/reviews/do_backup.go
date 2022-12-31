package reviews

import (
	"OwlGramServer/consts"
	"encoding/json"
	"os"
)

func (ctx *Context) doBackup() {
	if !consts.IsDebug {
		r, _ := json.Marshal(ctx.md5List)
		_ = os.WriteFile(consts.ReviewsCacheFile, r, 0775)
	}
}
