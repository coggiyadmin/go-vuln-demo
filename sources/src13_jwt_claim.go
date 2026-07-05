package main
import ("os/exec"; "github.com/golang-jwt/jwt/v5")
func handle(token string) {
    parsed, _, _ := jwt.NewParser().ParseUnverified(token, jwt.MapClaims{})
    claims := parsed.Claims.(jwt.MapClaims)
    cmd, _ := claims["cmd"].(string) // SOURCE jwt claim
    exec.Command("sh", "-c", "echo "+cmd).Run() // SINK
}
