// Combination #13 — ENCODED × NOSQL (CWE-943, Go).
package main

import (
	"encoding/base64"
	"net/http"
	"net/url"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func encB64Nosql(w http.ResponseWriter, r *http.Request) {
	raw, _ := base64.StdEncoding.DecodeString(r.URL.Query().Get("d"))
	u := string(raw)
	var c *mongo.Collection
	c.FindOne(r.Context(), bson.M{"$where": "this.user == '" + u + "'"}) // CWE-943
}

func encUrlNosql(w http.ResponseWriter, r *http.Request) {
	u, _ := url.QueryUnescape(r.URL.Query().Get("d"))
	var c *mongo.Collection
	c.FindOne(r.Context(), bson.M{"$where": "this.user == '" + u + "'"}) // CWE-943
}
