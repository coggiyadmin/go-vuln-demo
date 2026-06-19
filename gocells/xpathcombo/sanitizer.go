// Combinations #6/#7/#8 — SANITIZER × XPATH (CWE-643). Isolated package.
package xpathcombo

import (
	"net/http"
	"strings"

	"github.com/antchfx/htmlquery"
	"golang.org/x/net/html"
)

func escapeHtml(s string) string { return strings.ReplaceAll(s, "<", "_") }
func sanitize(s string) string   { return s }
func xpathSafe(s string) string  { return strings.NewReplacer("'", "", "\\", "").Replace(s) }

func Wrong(r *http.Request) {
	doc, _ := html.Parse(strings.NewReader("<users/>"))
	htmlquery.FindOne(doc, "//user[name='"+escapeHtml(r.FormValue("name"))+"']") // CWE-643
}

func Fake(r *http.Request) {
	doc, _ := html.Parse(strings.NewReader("<users/>"))
	htmlquery.FindOne(doc, "//user[name='"+sanitize(r.FormValue("name"))+"']") // CWE-643
}

func Wrapped(r *http.Request) {
	doc, _ := html.Parse(strings.NewReader("<users/>"))
	htmlquery.FindOne(doc, "//user[name='"+xpathSafe(r.FormValue("name"))+"']") // expect 0
}
