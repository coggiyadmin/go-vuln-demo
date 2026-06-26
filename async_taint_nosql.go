// Combination #4 — ASYNC taint × NOSQL (CWE-943, Go).
package main

import (
	"net/http"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func atNosqlHandler(w http.ResponseWriter, r *http.Request) {
	uid := r.FormValue("uid")
	ch := make(chan string, 1)
	go func() { ch <- uid }()
	v := <-ch
	var c *mongo.Collection
	c.FindOne(r.Context(), bson.M{"user": uid}) // CWE-943
}
