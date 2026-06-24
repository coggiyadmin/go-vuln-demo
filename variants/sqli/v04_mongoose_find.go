package main
import ("net/http"; "go.mongodb.org/mongo-driver/mongo")
func nosqlFind(w http.ResponseWriter, r *http.Request) {
    filter := r.FormValue("filter") // SOURCE — treated as query fragment
    _ = filter
    var c *mongo.Collection
    c.FindOne(r.Context(), map[string]interface{}{"user": filter}) // SINK CWE-943
}
