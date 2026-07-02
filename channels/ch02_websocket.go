package main
import "net/http"
func websocket(w http.ResponseWriter, r *http.Request) {
    msg := r.URL.Query().Get("msg") // SOURCE CH-02
    w.Write([]byte("<p>" + msg + "</p>"))
}
