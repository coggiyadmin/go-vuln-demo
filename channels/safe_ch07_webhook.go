package main
// TN — webhook field allowlisted, argv form.
import ("os/exec"; "regexp")
func onWebhook(repoName string) {
    if !regexp.MustCompile(`^[A-Za-z0-9_-]+$`).MatchString(repoName) {
        return
    }
    exec.Command("git", "clone", "/srv/"+repoName).Run()
}
