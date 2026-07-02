// TP — go install from user URL (CWE-506). CVE-2026-33634 B-tier class.
package intake

import (
	"net/http"
	"os/exec"
)

func InstallFromURL(w http.ResponseWriter, r *http.Request) {
	url := r.FormValue("url")
	exec.Command("go", "install", url+"@latest").Run() // SINK CWE-78 / supply_chain
}
