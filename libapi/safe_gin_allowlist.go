// Z6 benign mirror — outbound GET gated by package-level host allowlist (SSRF FP tax).
package libapi

import ("net/http"; "net/url")

var allowedHosts = map[string]bool{"api.internal.example.com": true}

func FetchSafe(raw string) (*http.Response, error) {
	u, err := url.Parse(raw)
	if err != nil {
		return nil, err
	}
	if allowedHosts[u.Hostname()] {
		return http.Get(raw)
	}
	return nil, nil
}
