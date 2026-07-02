package main

import "net/http"

func metaRefresh(w http.ResponseWriter, r *http.Request) {
    next := r.URL.Query().Get("next")
    w.Header().Set("Content-Type", "text/html")
    w.Write([]byte("<meta http-equiv=\"refresh\" content=\"0;url=" + next + "\">")) // CWE-601
}
