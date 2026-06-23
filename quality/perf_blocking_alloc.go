// Performance anti-pattern.
package quality

import (
	"io"
	"os"
)

func ProcessLog(path string) ([]byte, error) {
	return os.ReadFile(path)
}

func CopyAll(r io.Reader) ([]byte, error) {
	return io.ReadAll(r)
}
