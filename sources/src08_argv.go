package main
import ("os"; "os/exec")
func main() {
    v := ""
    if len(os.Args) > 1 { v = os.Args[1] } // SOURCE SRC-08
    exec.Command("sh", "-c", "echo "+v).Run()
}
