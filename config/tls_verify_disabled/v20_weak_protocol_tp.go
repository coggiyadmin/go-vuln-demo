package main
import "crypto/tls"
func cfg() *tls.Config {
    return &tls.Config{MinVersion: tls.VersionTLS10} // SINK weak min protocol
}
