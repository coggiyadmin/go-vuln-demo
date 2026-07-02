// Combination #11 — FAN-OUT × NOSQL (CWE-943, Go).
package main

import (
	"net/http"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func fanoutNosql(w http.ResponseWriter, r *http.Request) {
	u := r.URL.Query().Get("u")
	var c *mongo.Collection
	c.FindOne(r.Context(), bson.M{"$where": "this.a == '" + u + "'"}) // CWE-943
	c.FindOne(r.Context(), bson.M{"$where": "this.b == '" + u + "'"}) // CWE-943
}
