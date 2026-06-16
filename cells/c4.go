package cells
import ("fmt";"net/http";"os/exec")
func C4(r *http.Request){ c:=r.FormValue("c"); go func(){ exec.Command("sh","-c",fmt.Sprintf("echo %s",c)).Run() }() }
