package xpathcombo

import (
	"encoding/base64"
	"net/http"
	"net/url"
	"strings"

	"github.com/antchfx/htmlquery"
	"golang.org/x/net/html"
)

func EncodedB64(r *http.Request) {
	n, _ := base64.StdEncoding.DecodeString(r.FormValue("d"))
	doc, _ := html.Parse(strings.NewReader("<users/>"))
	htmlquery.FindOne(doc, "//user[name='"+string(n)+"']") // CWE-643
}

func EncodedURL(r *http.Request) {
	n, _ := url.QueryUnescape(r.FormValue("d"))
	doc, _ := html.Parse(strings.NewReader("<users/>"))
	htmlquery.FindOne(doc, "//user[name='"+n+"']") // CWE-643
}
