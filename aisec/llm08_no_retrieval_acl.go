// Vector/Embedding Weakness (OWASP LLM08) — retrieval with no tenant filter. TP.
package aisec

type Doc struct {
	Tenant string
	Text   string
}

var sharedIndex []Doc

func Retrieve(query string) []string {
	out := []string{}
	for _, d := range sharedIndex { // SINK (LLM08): no ACL filter on the shared store
		out = append(out, d.Text)
	}
	return out
}
