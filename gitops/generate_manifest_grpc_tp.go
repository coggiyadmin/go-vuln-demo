// GitOps TP — unauthenticated repo-server GenerateManifest (Synacktiv 2026-07).
// No auth on gRPC handler; user controls ManifestRequest → kustomize exec.
package gitops

import "context"

type ManifestRequest struct {
	RepoURL          string
	Revision         string
	KustomizeOptions *KustomizeOptions
}

type RepoService struct{}

// GenerateManifest — repository.RepoServerService/GenerateManifest (no auth middleware).
func (s *RepoService) GenerateManifest(ctx context.Context, q *ManifestRequest) error {
	// SOURCE — Parameter[1] in CodeQL model pack; reachable without authentication
	repoPath, err := cloneGitRepo(q.RepoURL, q.Revision)
	if err != nil {
		return err
	}
	return RunKustomizeBuild(q.KustomizeOptions, repoPath)
}

func cloneGitRepo(url, rev string) (string, error) {
	_ = rev
	return "/tmp/" + url, nil
}
