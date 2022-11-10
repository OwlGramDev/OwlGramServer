package updates

import (
	"OwlGramServer/consts"
	"OwlGramServer/updates/types"
	"encoding/json"
	"io/ioutil"
	"os"
)

func (ctx *Context) readUpdateInfo() {
	jsonFile, err := os.Open(consts.UpdateFileDescription)
	if err != nil {
		return
	}
	defer func(jsonFile *os.File) {
		_ = jsonFile.Close()
	}(jsonFile)
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return
	}
	var result types.UpdatesDescriptor
	err = json.Unmarshal(byteValue, &result)
	if err != nil {
		return
	}
	ctx.updateAppInfo(&result)
	ctx.UpdatesDescriptor = &result
}
