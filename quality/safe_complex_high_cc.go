// SAFE mirror.
package quality

func sign(n int) string {
	if n > 0 {
		return "p"
	}
	if n < 0 {
		return "n"
	}
	return "z"
}

func Classify(a, b int, kind string) string {
	if kind == "x" {
		return "x" + sign(a) + sign(b)
	}
	return "default"
}
