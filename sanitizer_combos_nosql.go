// Combinations #6/#7 — SANITIZER × NOSQL (CWE-943, Go).
package main

import (
	"net/http"
	"strings"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func wrongCtx(w http.ResponseWriter, r *http.Request) {
	u := strings.ReplaceAll(r.URL.Query().Get("user"), "<", "_")
	var c *mongo.Collection
	c.FindOne(r.Context(), bson.M{"$where": "this.a == '" + u + "'"}) // CWE-943
}

func fakeSan(w http.ResponseWriter, r *http.Request) {
	u := r.URL.Query().Get("user")
	var c *mongo.Collection
	c.FindOne(r.Context(), bson.M{"$where": "this.a == '" + u + "'"}) // CWE-943
}
