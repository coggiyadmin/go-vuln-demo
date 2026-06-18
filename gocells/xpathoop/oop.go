// Combination #5 — OOP OBJECT FLOW × XPATH INJECTION (CWE-643, Go). Isolated package.
// NO finding = FALSE NEGATIVE.
package xpathoop

import (
	"net/http"

	"github.com/antchfx/xmlquery"
)

type lookup struct {
	doc  *xmlquery.Node
	name string
}

func (l lookup) findDirect() []*xmlquery.Node {
	return xmlquery.Find(l.doc, "//user[name='"+l.name+"']") // XPath injection sink (CWE-643)
}

func Handle(doc *xmlquery.Node, r *http.Request) {
	l := lookup{doc: doc, name: r.FormValue("name")} // SOURCE -> struct field
	l.findDirect()
}
