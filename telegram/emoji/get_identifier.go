package emoji

import (
	"OwlGramServer/utilities"
	"github.com/gotd/td/tg"
	"strings"
)

func getIdentifier(docRaw *tg.MessageMediaDocument) string {
	document := docRaw.Document.(*tg.Document)
	attributes := document.Attributes
	var fileName string
	for _, x := range attributes {
		if utilities.InstanceOf(x, &tg.DocumentAttributeFilename{}) {
			fileName = x.(*tg.DocumentAttributeFilename).FileName
		}
	}
	if strings.Contains(fileName, ".") {
		fileName = fileName[:strings.LastIndex(fileName, ".")]
	}
	return fileName
}
