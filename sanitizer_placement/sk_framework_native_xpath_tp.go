package main
import ("html"; "net/http")
func skFrameworkNativeXpath(w http.ResponseWriter, r *http.Request) {
    q := html.EscapeString(r.URL.Query().Get('q'))
    _ = "//user[@name='" + q + "']"
}
