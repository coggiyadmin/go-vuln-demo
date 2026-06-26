// Combination #3 — LOOP-CARRIED TAINT × NOSQL (CWE-943, Go).
package main

import (
	"net/http"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func ltNosqlList(w http.ResponseWriter, r *http.Request) {
	in := r.Form["uid"]
	items := make([]string, 0, len(in))
	for _, x := range in { items = append(items, x) }
	uid := items[0]
	var c *mongo.Collection
	c.FindOne(r.Context(), bson.M{"user": uid}) // CWE-943
}
