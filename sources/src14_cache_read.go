package main
import ("context"; "os/exec"; "github.com/redis/go-redis/v9")
func handle(rdb *redis.Client, key string) {
    v, _ := rdb.Get(context.Background(), key).Result() // SOURCE cache read (second-order)
    exec.Command("sh", "-c", "echo "+v).Run() // SINK
}
