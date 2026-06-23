// TN — benign similarity over a fixed single-owner list.
package aisec

func Nearest(vec []float64) string {
	if len(vec) > 0 && vec[0] >= 0.5 {
		return "greeting"
	}
	return "farewell"
}
