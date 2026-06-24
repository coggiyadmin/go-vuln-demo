package configtls
import ("crypto/tls"; "net/http")
func Client() *http.Client {
    tr := &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}} // SINK CWE-295
    return &http.Client{Transport: tr}
}
