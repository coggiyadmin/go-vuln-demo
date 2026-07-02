// IL-62 — JSON env blob → format string (IL-5 config frontier).
package interop

import (
    "encoding/json"
    "fmt"
    "os"
)

func LookupEnvJsonFmt() {
    blob := os.Getenv("USER_FMT")
    if blob == "" {
        blob = "{}"
    }
    var m map[string]string
    _ = json.Unmarshal([]byte(blob), &m)
    _, _ = fmt.Fprintf(os.Stdout, m["fmt"], "guest")
}
