// SAFE mirror — argv array, no shell string.
package interop

import (
	"net/http"
	"os/exec"
)

func SafeShellInString(w http.ResponseWriter, r *http.Request) {
	arg := r.URL.Query().Get("arg")
	exec.Command("echo", "--", arg).Run()
}
