package main
// TN — MQTT payload allowlisted, argv form.
import ("os/exec"; "regexp")
func onMessage(payload []byte) {
    cmd := string(payload)
    if !regexp.MustCompile(`^[a-z_]+$`).MatchString(cmd) {
        return
    }
    exec.Command("handle", cmd).Run()
}
