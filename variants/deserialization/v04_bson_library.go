package main
import ("net/http"; "go.mongodb.org/mongo-driver/bson")
func bsonDecode(w http.ResponseWriter, r *http.Request) {
    var doc bson.M
    bson.Unmarshal([]byte(r.FormValue("blob")), &doc) // SINK CWE-502 bson
}
