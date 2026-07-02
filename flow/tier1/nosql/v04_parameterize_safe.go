// Phase-3 parameterize mirror — bound params / prepared API
package main
import ("net/http"; "go.mongodb.org/mongo-driver/bson"; "go.mongodb.org/mongo-driver/mongo")
func nosqlSafe(w http.ResponseWriter, r *http.Request) {
    user := r.FormValue("user")
    if user == "" { http.Error(w, "bad request", 400); return }
    var c *mongo.Collection
    c.FindOne(r.Context(), bson.M{"user": user})
}
