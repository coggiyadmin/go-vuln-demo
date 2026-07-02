// Combination #5 — OOP × NOSQL (CWE-943, Go).
package main

import (
	"net/http"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Query struct{ expr string }

func (q *Query) findDirect(c *mongo.Collection) {
	c.FindOne(nil, bson.M{"$where": "this.user == '" + q.expr + "'"}) // CWE-943
}

func oopNosql(w http.ResponseWriter, r *http.Request) {
	q := &Query{expr: r.URL.Query().Get("user")}
	var c *mongo.Collection
	q.findDirect(c)
}
