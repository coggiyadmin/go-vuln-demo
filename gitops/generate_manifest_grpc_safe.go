// SAFE — authenticated gRPC handler; rejects unauthenticated GenerateManifest calls.
package gitops

import (
	"context"
	"errors"
)

type authContextKey struct{}

func (s *RepoService) GenerateManifestSafe(ctx context.Context, q *ManifestRequest) error {
	if ctx.Value(authContextKey{}) == nil {
		return errors.New("unauthenticated")
	}
	repoPath, err := cloneGitRepo(q.RepoURL, q.Revision)
	if err != nil {
		return err
	}
	return RunKustomizeBuildSafe(q.KustomizeOptions, repoPath)
}
