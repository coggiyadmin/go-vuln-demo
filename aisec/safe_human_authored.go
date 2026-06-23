// SAFE mirror.
package aisec

func SumPositive(values []int) int {
	sum := 0
	for _, v := range values {
		if v > 0 {
			sum += v
		}
	}
	return sum
}
