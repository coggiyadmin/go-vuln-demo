// Safe mirror — pinned module path only.
package intake

import (
	"net/http"
	"os/exec"
)

func InstallSafe(w http.ResponseWriter, r *http.Request) {
	const mod = "example.com/safe/mod@v1.0.0"
	exec.Command("go", "install", mod).Run()
}
