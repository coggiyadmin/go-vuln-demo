// SAFE — constant format only.
package interop

import "fmt"

func LookupEnvJsonFmtSafe() {
    _ = fmt.Sprintf("%s", "guest")
}
