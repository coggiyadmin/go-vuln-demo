package main
import ("os"; "os/exec")
func main() {
    v := os.Getenv("TARGET") // SOURCE SRC-07
    exec.Command("sh", "-c", "echo "+v).Run()
}
