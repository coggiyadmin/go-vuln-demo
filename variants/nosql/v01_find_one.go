package main
import ("net/http"; "go.mongodb.org/mongo-driver/bson"; "go.mongodb.org/mongo-driver/mongo")
func nosqlLogin(w http.ResponseWriter, r *http.Request) {
    user := r.FormValue("user")
    var c *mongo.Collection
    c.FindOne(r.Context(), bson.M{"user": user}) // SINK CWE-943
}
