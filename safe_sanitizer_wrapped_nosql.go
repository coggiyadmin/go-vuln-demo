// c08 SAFE — validated key wrapper × NoSQL. Expect clean.
package main

import (
	"net/http"
	"regexp"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var keyRe = regexp.MustCompile(`^[a-zA-Z0-9_-]+$`)

func wrappedNosql(w http.ResponseWriter, r *http.Request) {
	k := r.URL.Query().Get("k")
	if !keyRe.MatchString(k) {
		return
	}
	var c *mongo.Collection
	c.FindOne(r.Context(), bson.M{"user": k})
}
