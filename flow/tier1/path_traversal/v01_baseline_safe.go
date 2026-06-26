package main
import ("path/filepath"; "net/http"; "os")
func pathTravSafe(w http.ResponseWriter, r *http.Request) {
    root := "/data"; p := r.URL.Query().Get("p")
    full := filepath.Clean(filepath.Join(root, p))
    if full[:len(root)] != root { http.Error(w, "forbidden", 403); return }
    os.ReadFile(full)
}
