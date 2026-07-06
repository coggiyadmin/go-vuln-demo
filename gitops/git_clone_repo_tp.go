// GitOps TP — attacker-controlled Git URL fetched by repo-server (SSRF / supply-chain).
package gitops

import "os/exec"

func FetchApplicationManifests(repoURL string) error {
	workDir := "/tmp/repos/app"
	// SOURCE — ManifestRequest.Repo.Repo from unauthenticated gRPC
	// SINK CWE-918 — internal/metadata URL fetch
	return exec.Command("git", "clone", "--depth", "1", repoURL, workDir).Run()
}
