package xpathxf

import (
	"net/http"

	"github.com/antchfx/xmlquery"
)

func Handle(doc *xmlquery.Node, r *http.Request) { Run(doc, r.FormValue("name")) } // SOURCE -> cross-file sink (CWE-643)
