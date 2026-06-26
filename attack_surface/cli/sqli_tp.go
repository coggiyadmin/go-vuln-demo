package main
import "os/exec"
import "os"
func main() {
    v := os.Args[1]
    exec.Command("sh", "-c", "grep "+v).Run()
}
