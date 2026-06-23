// McCabe HIGH CC.
package quality

func Classify(a, b, c, d int, kind string) string {
	if kind == "x" {
		if a > 0 {
			if b > 0 {
				if c > 0 {
					if d > 0 {
						return "xppp"
					}
					return "xppn"
				}
			}
		}
	}
	if kind == "y" {
		if a > 0 && b > 0 {
			return "y1"
		}
	}
	return "default"
}
