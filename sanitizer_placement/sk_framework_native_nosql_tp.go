package main
import ("encoding/json"; "fmt"; "net/http")
func skFrameworkNativeNosql(w http.ResponseWriter, r *http.Request) {
    var m map[string]interface{}
    _ = json.Unmarshal([]byte(r.FormValue("user")), &m)
    _ = fmt.Sprintf(`{"user":%v}`, m)
}
