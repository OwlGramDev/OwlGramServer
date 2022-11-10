package reviews

import (
	"OwlGramServer/consts"
	"encoding/json"
	"os"
)

func (ctx *Context) readBackup() {
	r, err := os.ReadFile(consts.ReviewsCacheFile)
	if err == nil {
		var recovery []string
		_ = json.Unmarshal(r, &recovery)
		ctx.md5List = recovery
	}
}
