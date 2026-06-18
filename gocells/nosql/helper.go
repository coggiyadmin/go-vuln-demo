package nosqlxf

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func Run(coll *mongo.Collection, expr string) (*mongo.Cursor, error) {
	return coll.Find(context.TODO(), bson.M{"$where": "this.user == '" + expr + "'"}) // SINK CWE-943
}
