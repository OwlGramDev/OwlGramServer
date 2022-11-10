package tg_checker

import (
	"OwlGramServer/consts"
	"encoding/json"
	"os"
)

func (tg *Context) doBackup() {
	r, _ := json.Marshal(tg.StatusDC)
	_ = os.WriteFile(consts.DcCheckerCacheFile, r, 0775)
}
