package nosqlxf

import (
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
)

func Handle(coll *mongo.Collection, r *http.Request) { Run(coll, r.FormValue("user")) } // SOURCE -> cross-file sink (CWE-943)
