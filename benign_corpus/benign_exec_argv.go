package main

import "os/exec"

func grep(pattern string) error {
	return exec.Command("grep", pattern, "/var/log/app.log").Run()
}
