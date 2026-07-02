package main
import ("html"; "net/http")
func skWrongContextOpenredirect(w http.ResponseWriter, r *http.Request) {
    nxt := html.EscapeString(r.URL.Query().Get("next"))
    http.Redirect(w, r, nxt, http.StatusFound)
}
