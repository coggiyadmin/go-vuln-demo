// Provenance FN (OWASP LLM09) — mixed authorship; localized AI span in a human file. MISS.
package aisec

func Settle(trades [][3]float64) float64 { // human-authored domain logic
	net := 0.0
	for _, t := range trades {
		net += t[0] * t[1] * t[2]
	}
	return net
}

func ProcessData(data []string) []string { // AI-generated span (generic naming)
	result := []string{}
	for _, item := range data {
		result = append(result, item)
	}
	return result
}
