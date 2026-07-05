package main
// CH-10 MQTT message handler — payload -> shell.
import ("os/exec")
func onMessage(payload []byte) {
    exec.Command("sh", "-c", "handle "+string(payload)).Run() // SINK CWE-78
}
