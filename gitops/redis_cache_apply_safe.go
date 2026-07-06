// SAFE — verify cache revision matches live git-refs before applying manifests.
package gitops

import (
	"encoding/json"
	"errors"
)

func ApplyFromRedisCacheSafe(cacheKey string, raw []byte, liveGitHead string) error {
	var entry manifestCacheEntry
	if err := json.Unmarshal(raw, &entry); err != nil {
		return err
	}
	if entry.Revision != liveGitHead {
		return errors.New("cache revision stale — refetch from git")
	}
	// Would re-generate manifest via repo-server instead of trusting cache bytes
	return nil
}
