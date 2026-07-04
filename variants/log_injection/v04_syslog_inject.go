package main
import ("log"; "net/http"; "os")
func syslogLog(w http.ResponseWriter, r *http.Request) {
    user := r.FormValue("user")
    f, _ := os.OpenFile("/var/log/app.log", os.O_APPEND|os.O_WRONLY, 0644)
    log.SetOutput(f)
    log.Printf("user=%s", user) // SINK CWE-117
}
