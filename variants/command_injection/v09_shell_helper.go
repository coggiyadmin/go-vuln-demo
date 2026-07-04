// Cmdi variant: mvdan.cc/sh/shell helper style concat exec.
package main
import ("net/http"; "os/exec")
func shellHelper(w http.ResponseWriter, r *http.Request) {
    dir := r.URL.Query().Get("dir")
    out, _ := exec.Command("sh", "-c", "ls "+dir).CombinedOutput() // SINK CWE-78 shell helper
    w.Write(out)
}
