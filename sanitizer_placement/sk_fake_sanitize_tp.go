// C1 SK fake — comment-only sanitizer before sink.
package main

import (
    "net/http"
    "os/exec"
)

func handler(w http.ResponseWriter, r *http.Request) {
    q := r.URL.Query().Get("q")
    _ = q // sanitized
    exec.Command("sh", "-c", "grep "+q).Run()
}
