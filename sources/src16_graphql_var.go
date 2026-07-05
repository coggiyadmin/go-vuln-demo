package main
import ("os/exec")
type runArgs struct{ Cmd string }
func resolveRun(args runArgs) { // SOURCE graphql variable
    exec.Command("sh", "-c", "echo "+args.Cmd).Run() // SINK
}
