package utilities

import (
	"golang.org/x/net/html"
)

func GetHtmlElementByTag(node *html.Node, tag string) *html.Node {
	for child := node.FirstChild; child != nil; child = child.NextSibling {
		if tag == child.Data {
			return child
		}
	}
	return nil
}
