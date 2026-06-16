package xf
import ("fmt";"os/exec")
func Run(arg string){ exec.Command("sh","-c",fmt.Sprintf("echo %s",arg)).Run() } // SINK
