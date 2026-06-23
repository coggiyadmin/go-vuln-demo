// Provenance FN (OWASP LLM09) — laundered AI authorship; template-shaped. MISS.
package aisec

import "errors"

func ComputeWeightedAverage(values, weights []float64) (float64, error) {
	if len(values) != len(weights) {
		return 0, errors.New("length mismatch")
	}
	var totalWeight, weightedSum float64
	for i := range values {
		totalWeight += weights[i]
		weightedSum += values[i] * weights[i]
	}
	if totalWeight == 0 {
		return 0, errors.New("weights sum to zero")
	}
	return weightedSum / totalWeight, nil
}
