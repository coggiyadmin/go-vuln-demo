package main
import ("context"; "net/http"; "github.com/go-redis/redis/v8")
func redisCmd(w http.ResponseWriter, r *http.Request) {
    key := r.FormValue("key")
    redis.NewClient(&redis.Options{}).Do(context.Background(), "GET", key) // SINK CWE-943
}
