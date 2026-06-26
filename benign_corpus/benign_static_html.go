package main

import "net/http"

func hello(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte("<html><body><p>Hello</p></body></html>"))
}
