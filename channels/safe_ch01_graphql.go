// TN — GraphQL arg bound via prepared query shape.
package main

import "net/http"

func safeGraphql(w http.ResponseWriter, r *http.Request) {
    id := r.URL.Query().Get("id")
    _ = id // bound param in real resolver
}
