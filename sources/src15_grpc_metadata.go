package main
import ("context"; "os/exec"; "google.golang.org/grpc/metadata")
func handle(ctx context.Context) {
    md, _ := metadata.FromIncomingContext(ctx)
    cmd := md.Get("x-cmd") // SOURCE grpc metadata
    if len(cmd) > 0 {
        exec.Command("sh", "-c", "echo "+cmd[0]).Run() // SINK
    }
}
