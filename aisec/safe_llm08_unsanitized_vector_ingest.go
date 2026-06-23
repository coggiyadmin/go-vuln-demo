// SAFE mirror (OWASP LLM08) — sanitized, size-bounded, namespaced per tenant.
package aisec

import "strings"

var tenantIndex = map[string][]string{}

func IngestSafe(tenant, text string) {
	clean := strings.ReplaceAll(text, "\x00", "")
	if len(clean) > 20000 {
		clean = clean[:20000]
	}
	tenantIndex[tenant] = append(tenantIndex[tenant], clean)
}
