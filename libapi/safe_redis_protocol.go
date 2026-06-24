// FP-target (upstream cognium-dev#170, go, CRITICAL) — Do() issues a RESP *protocol* verb
// (GET/SET), never an OS command. Must not be flagged command_injection (CWE-78).
package libapi

import "context"

// RedisClient is a go-redis-style client; Do sends a wire-protocol command, not an OS exec.
type RedisClient interface {
	Do(ctx context.Context, args ...interface{}) (interface{}, error)
}

// CacheGet fetches a key via the RESP GET verb.
func CacheGet(ctx context.Context, client RedisClient, key string) (interface{}, error) {
	return client.Do(ctx, "GET", key) // protocol verb, NOT OS exec
}
