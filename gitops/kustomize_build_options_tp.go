// GitOps TP — Argo CD repo-server kustomize.go shape (Synacktiv 2026-07).
// BuildOptions from ManifestRequest flows into strings.Fields → exec.Command.
package gitops

import (
	"os/exec"
	"strings"
)

type KustomizeOptions struct {
	BuildOptions string // SOURCE — gRPC ManifestRequest.KustomizeOptions
	BinaryPath   string
}

type kustomizeRunner struct {
	repoPath   string
	binaryPath string
}

func (k *kustomizeRunner) binary() string {
	if k.binaryPath != "" {
		return k.binaryPath
	}
	return "kustomize"
}

func parseBuildOptions(path, buildOptions string) []string {
	return append([]string{"build", path}, strings.Fields(buildOptions)...)
}

func RunKustomizeBuild(opts *KustomizeOptions, repoPath string) error {
	k := &kustomizeRunner{repoPath: repoPath}
	if opts != nil {
		k.binaryPath = opts.BinaryPath
		params := parseBuildOptions(repoPath, opts.BuildOptions)
		// SINK CWE-78 — e.g. --enable-helm --helm-command ./pwn.sh
		return exec.Command(k.binary(), params...).Run()
	}
	return exec.Command(k.binary(), "build", repoPath).Run()
}
