// GitOps TP — Redis cache poisoning (Synacktiv / Cycode 2024 class).
// Controller trusts gzip JSON in mfst|* keys without re-fetching Git.
package gitops

import (
	"encoding/json"
	"os/exec"
	"strings"
)

type manifestCacheEntry struct {
	Manifests []string `json:"manifests"`
	Revision  string   `json:"revision"`
}

// ApplyFromRedisCache — sync loop reads attacker-poisoned cache entry.
func ApplyFromRedisCache(cacheKey string, raw []byte) error {
	// SOURCE — Redis GET; attacker writes mfst|... and downgrades git-refs|...
	var entry manifestCacheEntry
	if err := json.Unmarshal(raw, &entry); err != nil {
		return err
	}
	for _, manifestJSON := range entry.Manifests {
		// SINK — arbitrary cluster manifest applied from untrusted cache
		cmd := exec.Command("kubectl", "apply", "-f", "-")
		cmd.Stdin = strings.NewReader(manifestJSON)
		if err := cmd.Run(); err != nil {
			return err
		}
	}
	return nil
}
