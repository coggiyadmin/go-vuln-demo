package main
import ("encoding/json"; "net/http")
func mongoAgg(w http.ResponseWriter, r *http.Request) {
    stage := r.FormValue("stage")
    _ = json.RawMessage(`[{"$match":` + stage + `}]`) // SINK CWE-943
}
