// Vector/Embedding FN (OWASP LLM08) — shared cache keyed only by query, not tenant. MISS.
package aisec

var cache = map[string][]string{}

func Retrieve(tenant, query string) []string {
	if hit, ok := cache[query]; ok { // SINK (LLM08 FN): cache key omits tenant
		return hit
	}
	res := []string{tenant + ":" + query}
	cache[query] = res
	return res
}
