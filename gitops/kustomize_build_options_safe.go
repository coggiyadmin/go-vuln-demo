// SAFE — allowlisted kustomize build flags only (no --helm-command / arbitrary tokens).
package gitops

import (
	"os/exec"
	"strings"
)

var allowedKustomizeFlags = map[string]bool{
	"--load-restrictor": true, "--reorder": true, "--enable-alpha-plugins": true,
}

func RunKustomizeBuildSafe(opts *KustomizeOptions, repoPath string) error {
	if opts == nil || opts.BuildOptions == "" {
		return exec.Command("kustomize", "build", repoPath).Run()
	}
	tokens := strings.Fields(opts.BuildOptions)
	safe := make([]string, 0, len(tokens))
	for i := 0; i < len(tokens); i++ {
		t := tokens[i]
		if !allowedKustomizeFlags[t] {
			continue
		}
		safe = append(safe, t)
		if i+1 < len(tokens) && !strings.HasPrefix(tokens[i+1], "-") {
			safe = append(safe, tokens[i+1])
			i++
		}
	}
	params := append([]string{"build", repoPath}, safe...)
	return exec.Command("kustomize", params...).Run()
}
