package main
import "crypto/tls"
func cfg() *tls.Config {
    return &tls.Config{MinVersion: tls.VersionTLS12} // TLS 1.2 minimum
}
