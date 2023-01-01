package emoji

import (
	"bytes"
	"github.com/gotd/td/telegram/downloader"
	"github.com/gotd/td/tg"
)

func (c *Context) downloadBytes(docRaw *tg.MessageMediaDocument) ([]byte, error) {
	document := docRaw.Document.(*tg.Document)
	d := downloader.NewDownloader()
	loc := document.AsInputDocumentFileLocation()
	var file bytes.Buffer
	_, err := d.Download(c.client, loc).Stream(c.context, &file)
	if err != nil {
		return nil, err
	}
	return file.Bytes(), nil
}
