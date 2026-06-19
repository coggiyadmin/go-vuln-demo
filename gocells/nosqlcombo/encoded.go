package nosqlcombo

import (
	"context"
	"encoding/base64"
	"net/http"
	"net/url"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func EncodedB64(coll *mongo.Collection, r *http.Request) {
	e, _ := base64.StdEncoding.DecodeString(r.FormValue("d"))
	coll.Find(context.TODO(), bson.M{"$where": "this.user == '" + string(e) + "'"}) // CWE-943
}

func EncodedURL(coll *mongo.Collection, r *http.Request) {
	e, _ := url.QueryUnescape(r.FormValue("d"))
	coll.Find(context.TODO(), bson.M{"$where": "this.user == '" + e + "'"}) // CWE-943
}
