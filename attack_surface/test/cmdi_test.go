package attacksurface
import ("os"; "os/exec"; "testing")
func TestRun(t *testing.T) {
    q := os.Getenv("Q")
    exec.Command("sh", "-c", "grep "+q).Run() // SINK — test file
}
