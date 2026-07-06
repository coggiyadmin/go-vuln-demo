// SAFE — host allowlist before git clone.
package gitops

import (
	"errors"
	"net/url"
	"os/exec"
	"strings"
)

var allowedGitHosts = map[string]bool{
	"github.com": true, "gitlab.com": true,
}

func FetchApplicationManifestsSafe(repoURL string) error {
	u, err := url.Parse(repoURL)
	if err != nil || u.Scheme != "https" {
		return errors.New("invalid repo url")
	}
	host := strings.ToLower(u.Hostname())
	if !allowedGitHosts[host] {
		return errors.New("repo host not allowlisted")
	}
	return exec.Command("git", "clone", "--depth", "1", repoURL, "/tmp/repos/app").Run()
}
