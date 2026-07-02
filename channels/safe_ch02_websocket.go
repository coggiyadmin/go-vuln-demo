// TN — websocket payload HTML-escaped before echo.
package main

import (
    "html"
    "net/http"
)

func safeWebsocket(w http.ResponseWriter, r *http.Request) {
    msg := r.URL.Query().Get("msg")
    w.Write([]byte(html.EscapeString(msg)))
}
