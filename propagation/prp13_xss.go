package main
import ("encoding/json"; "net/http")
func handleRoundtripxss(w http.ResponseWriter, r *http.Request) {
    s := r.FormValue("q") // SOURCE
    b, _ := json.Marshal(map[string]string{"v": s}); var m map[string]string; json.Unmarshal(b, &m); t := m["v"] // round-trip
    w.Write([]byte("<p>" + t + "</p>")) // SINK
}
