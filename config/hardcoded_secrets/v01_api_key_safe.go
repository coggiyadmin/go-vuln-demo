package configsecret
import "os"
func APIKey() string { return os.Getenv("API_KEY") }
func DBPassword() string { return os.Getenv("DB_PASSWORD") }
