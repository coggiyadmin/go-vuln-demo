// Application entry-point SSRF (TP) — contrast with lib helper.
package profile

import ("io"; "net/http")

func ServerFetch(w http.ResponseWriter, r *http.Request) {
    url := r.URL.Query().Get("url") // SOURCE
    resp, _ := http.Get(url)        // SINK CWE-918
    io.Copy(w, resp.Body)
}
