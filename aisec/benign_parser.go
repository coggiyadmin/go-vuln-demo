// TN — benign parser; trims and drops blank lines.
package aisec

import "strings"

func ParseLines(lines []string) []string {
	out := []string{}
	for _, l := range lines {
		if t := strings.TrimSpace(l); t != "" {
			out = append(out, t)
		}
	}
	return out
}
