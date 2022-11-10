package updates

import (
	"html"
	"regexp"
	"strings"
)

func getReadableResult(postHtml string) string {
	r1 := regexp.MustCompile(`<i\sclass="emoji"\sstyle="background-image:url\(.*?\)">(.*?)</i>`)
	r2 := regexp.MustCompile(`<tg-emoji\semoji-id=".*?">(.*?)</tg-emoji>`)
	r3 := regexp.MustCompile(`target=".*?" rel=".*?" onclick=".*?"`)
	r4 := regexp.MustCompile(`<b>🔄.*?APK.*>`)
	r5 := regexp.MustCompile(`<b>📦.*?APK.*>`)
	r6 := regexp.MustCompile(`target=".*?" rel=".*?"`)
	r7 := regexp.MustCompile(`" >`)

	postHtml = r1.ReplaceAllString(postHtml, "$1")
	postHtml = r2.ReplaceAllString(postHtml, "$1")
	postHtml = strings.ReplaceAll(postHtml, "<br/>", "\n")
	postHtml = r3.ReplaceAllString(postHtml, "")
	postHtml = r4.ReplaceAllString(postHtml, "")
	postHtml = r5.ReplaceAllString(postHtml, "")
	postHtml = r6.ReplaceAllString(postHtml, "")
	postHtml = r7.ReplaceAllString(postHtml, "\">")
	postHtml = strings.ReplaceAll(postHtml, "<code>", "<tt>")
	postHtml = strings.ReplaceAll(postHtml, "</code>", "</tt>")
	postHtml = html.UnescapeString(postHtml)
	postHtml = strings.Trim(postHtml, "\n")
	return postHtml
}
