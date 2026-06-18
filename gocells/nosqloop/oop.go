// Combination #5 — OOP OBJECT FLOW × NoSQL INJECTION (CWE-943, Go). Isolated package.
// NO finding = FALSE NEGATIVE.
package nosqloop

import (
	"context"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type query struct {
	coll *mongo.Collection
	expr string
}

func (q query) findDirect() (*mongo.Cursor, error) {
	return q.coll.Find(context.TODO(), bson.M{"$where": "this.user == '" + q.expr + "'"}) // NoSQL injection (CWE-943)
}

func Handle(coll *mongo.Collection, r *http.Request) {
	q := query{coll: coll, expr: r.FormValue("user")} // SOURCE -> struct field
	q.findDirect()
}
