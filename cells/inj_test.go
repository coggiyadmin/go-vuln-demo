package cells
import ("fmt";"net/http";"os/exec")
func Inj(r *http.Request){ c:=r.FormValue("c"); s:=fmt.Sprintf("echo %s",c); exec.Command("sh","-c",s).Run() }
