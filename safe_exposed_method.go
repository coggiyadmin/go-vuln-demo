// SAFE mirror — exposed_method.go; the privileged action requires an authenticated admin.
package main

import (
	"net/http"
	"os"
	"os/exec"
)

func safeExposedMaintenance(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("X-Admin-Token") != os.Getenv("ADMIN_TOKEN") { // auth gate
		http.Error(w, "forbidden", http.StatusForbidden)
		return
	}
	exec.Command("/opt/app/bin/cleanup.sh").Run()
}
