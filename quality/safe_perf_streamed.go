// SAFE mirror — bounded buffer.
package quality

import (
	"bufio"
	"os"
)

func ProcessLog(path string, maxLines int) ([]string, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	sc := bufio.NewScanner(f)
	lines := []string{}
	for sc.Scan() && len(lines) < maxLines {
		lines = append(lines, sc.Text())
	}
	return lines, sc.Err()
}
