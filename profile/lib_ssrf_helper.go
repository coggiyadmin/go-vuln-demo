// Library-API URL param — NOT an HTTP attacker source (#128).
package profile

import "net/http"

func FetchURL(client *http.Client, targetURL string) (*http.Response, error) {
    return client.Get(targetURL) // FP-target — library param
}
