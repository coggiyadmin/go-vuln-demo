package main
import ("context"; "net/http")
import "go.mongodb.org/mongo-driver/bson"
import "go.mongodb.org/mongo-driver/mongo"
func query(w http.ResponseWriter, r *http.Request, coll *mongo.Collection) {
    name := r.FormValue("name") // SOURCE
    coll.Find(context.TODO(), bson.M{"$where": "this.name == '" + name + "'"}) // SINK CWE-943 mongo-go $where
}
