package main
// CH-07 webhook — payload field -> shell.
import ("os/exec")
func onWebhook(repoName string) {
    exec.Command("sh", "-c", "git clone /srv/"+repoName).Run() // SINK CWE-78
}
