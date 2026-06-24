// FP-target (upstream cognium-dev#162/#140) — CLI profile. A dev tool that execs an argv arg.
package main

import (
	"os"
	"os/exec"
)

func main() {
	exec.Command(os.Args[1]).Run() // dev-CLI, operator-controlled
}
