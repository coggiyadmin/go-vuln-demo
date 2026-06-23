// SAFE mirror (OWASP LLM06) — fs tool jailed to a workspace with a resolved-prefix check.
package aisec

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
)

var workspace, _ = filepath.Abs("/var/app/workspace")

func ReadToolSafe(rel string) (string, error) {
	target := filepath.Join(workspace, rel)
	if target != workspace && !strings.HasPrefix(target, workspace+string(os.PathSeparator)) {
		return "", errors.New("escapes jail")
	}
	b, err := os.ReadFile(target)
	return string(b), err
}
