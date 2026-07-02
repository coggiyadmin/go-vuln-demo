// IL-47 — CSV config field → exec (IL-5 config frontier).
package interop

import (
    "encoding/csv"
    "os/exec"
    "strings"
)

func RunFromCSV(raw string) error {
    rows, _ := csv.NewReader(strings.NewReader(raw)).ReadAll()
    cmd := rows[0][0] // SOURCE
    return exec.Command("sh", "-c", cmd).Run() // SINK CWE-78
}
