package xpathcombo

import (
	"net/http"
	"strings"

	"github.com/antchfx/htmlquery"
	"golang.org/x/net/html"
)

func Fanout(r *http.Request) {
	n := r.FormValue("n")
	doc, _ := html.Parse(strings.NewReader("<users/>"))
	htmlquery.FindOne(doc, "//user[name='"+n+"']") // CWE-643
	htmlquery.FindOne(doc, "//account[name='"+n+"']") // CWE-643
	htmlquery.FindOne(doc, "//*[@id='"+n+"']") // CWE-643
}
