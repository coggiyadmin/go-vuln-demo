package main
import ("net/http"; "os/exec")
func grpc(w http.ResponseWriter, r *http.Request) {
    p := r.URL.Query().Get("payload") // SOURCE CH-03
    exec.Command("sh", "-c", "echo "+p).Run()
}
