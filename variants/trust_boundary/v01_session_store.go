// trust_boundary variant: gorilla session Values map.
package main
import (
    "net/http"
    "github.com/gorilla/sessions"
)
var store = sessions.NewCookieStore([]byte("dev"))
func trustSession(w http.ResponseWriter, r *http.Request) {
    sess, _ := store.Get(r, "sid")
    sess.Values["role"] = r.FormValue("role") // SINK CWE-501
    sess.Save(r, w)
}
