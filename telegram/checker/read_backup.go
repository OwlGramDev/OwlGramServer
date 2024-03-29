package checker

import (
	"OwlGramServer/consts"
	"OwlGramServer/telegram/checker/types"
	"encoding/json"
	"os"
)

func (tg *Context) readBackup() {
	r, err := os.ReadFile(consts.DcCheckerCacheFile)
	if err == nil {
		var recovery []types.DCStatus
		_ = json.Unmarshal(r, &recovery)
		tg.StatusDC = recovery
	}
}
