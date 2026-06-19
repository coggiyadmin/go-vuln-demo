// Combinations #6/#7/#8 — SANITIZER × NoSQL (CWE-943). Isolated package.
package nosqlcombo

import (
	"context"
	"net/http"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func escapeHtml(s string) string { return strings.ReplaceAll(s, "<", "_") }
func sanitize(s string) string   { return s }

func Wrong(coll *mongo.Collection, r *http.Request) {
	coll.Find(context.TODO(), bson.M{"$where": "this.user == '" + escapeHtml(r.FormValue("user")) + "'"}) // CWE-943
}

func Fake(coll *mongo.Collection, r *http.Request) {
	coll.Find(context.TODO(), bson.M{"$where": "this.user == '" + sanitize(r.FormValue("user")) + "'"}) // CWE-943
}

func Wrapped(coll *mongo.Collection, r *http.Request) {
	coll.Find(context.TODO(), bson.M{"user": r.FormValue("user")}) // expect 0
}
