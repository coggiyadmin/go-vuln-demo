package xpathxf

import "github.com/antchfx/xmlquery"

func Run(doc *xmlquery.Node, name string) []*xmlquery.Node {
	return xmlquery.Find(doc, "//user[name='"+name+"']") // SINK CWE-643
}
