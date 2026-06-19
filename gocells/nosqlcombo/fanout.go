package nosqlcombo

import (
	"context"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func Fanout(coll *mongo.Collection, r *http.Request) {
	u := r.FormValue("u")
	coll.Find(context.TODO(), bson.M{"$where": "this.a == '" + u + "'"}) // CWE-943
	coll.Find(context.TODO(), bson.M{"$where": "this.b == '" + u + "'"}) // CWE-943
	coll.Find(context.TODO(), bson.M{"$where": "this.c == '" + u + "'"}) // CWE-943
}
