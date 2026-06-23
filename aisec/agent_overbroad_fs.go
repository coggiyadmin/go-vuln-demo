// Excessive Agency (OWASP LLM06) — over-broad fs tool rooted at "/", no jail.
package aisec

import (
	"os"
	"path/filepath"
)

const fsRoot = "/" // SINK (LLM06): whole filesystem in scope

func ReadTool(rel string) (string, error) {
	b, err := os.ReadFile(filepath.Join(fsRoot, rel))
	return string(b), err
}
