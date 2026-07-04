package main
import ("github.com/aws/aws-sdk-go/service/dynamodb"; "net/http")
func dynamoGet(w http.ResponseWriter, r *http.Request) {
    key := r.URL.Query().Get("key")
    _ = &dynamodb.GetItemInput{TableName: strPtr("users"), Key: map[string]*dynamodb.AttributeValue{
        "id": {S: &key},
    }} // SINK CWE-943
}
func strPtr(s string) *string { return &s }
