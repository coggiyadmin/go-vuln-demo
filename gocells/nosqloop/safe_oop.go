// SAFE mirror — nosqloop; value-bound equality match, no $where / operator injection.
// Expect 0 security findings.
package nosqloop

import (
	"context"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type safeQuery struct {
	coll *mongo.Collection
	expr string
}

func (q safeQuery) findDirect() (*mongo.Cursor, error) {
	return q.coll.Find(context.TODO(), bson.M{"user": q.expr}) // value-bound equality
}

func SafeHandle(coll *mongo.Collection, r *http.Request) {
	q := safeQuery{coll: coll, expr: r.FormValue("user")}
	q.findDirect()
}
