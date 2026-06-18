// SAFE mirror — xpathoop; input constrained to alphanumerics before building the query.
// Expect 0 security findings.
package xpathoop

import (
	"net/http"
	"regexp"

	"github.com/antchfx/xmlquery"
)

var alnum = regexp.MustCompile(`^[A-Za-z0-9]+$`)

type safeLookup struct {
	doc  *xmlquery.Node
	name string
}

func (l safeLookup) findDirect() []*xmlquery.Node {
	if !alnum.MatchString(l.name) {
		return nil
	}
	return xmlquery.Find(l.doc, "//user[name='"+l.name+"']") // input constrained to alnum
}

func SafeHandle(doc *xmlquery.Node, r *http.Request) {
	l := safeLookup{doc: doc, name: r.FormValue("name")}
	l.findDirect()
}
