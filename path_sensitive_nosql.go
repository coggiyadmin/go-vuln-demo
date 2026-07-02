// Combination #2 — PATH-SENSITIVITY × NOSQL (CWE-943, Go).
package main

import (
	"net/http"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func psNosqlNeg(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query().Get("q")
	var c *mongo.Collection
	if q == "safe" {
		return
	}
	c.FindOne(r.Context(), bson.M{"name": q}) // CWE-943
}

func psNosqlOneBranch(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query().Get("q")
	if false {
		q = "safe_literal"
	}
	var c *mongo.Collection
	c.FindOne(r.Context(), bson.M{"name": q}) // CWE-943
}
