package main
import "os/exec"
func PublicRun(user string) error {
    return exec.Command("sh", "-c", "grep "+user).Run() // TN no HTTP entry
}
