// SAFE mirror (OWASP LLM08) — retrieval scoped to the caller's tenant first.
package aisec

var scopedIndex []Doc

func RetrieveSafe(tenant, query string) []string {
	out := []string{}
	for _, d := range scopedIndex {
		if d.Tenant == tenant {
			out = append(out, d.Text)
		}
	}
	return out
}
