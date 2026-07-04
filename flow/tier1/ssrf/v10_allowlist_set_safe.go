// v10 allowlist-set — prefix slice gate before outbound GET.
package main

import ("net/http"; "strings")

var allowedPrefixes = []string{"https://cdn.example.com", "https://api.example.com"}

func fetch(raw string) (*http.Response, error) {
	for _, p := range allowedPrefixes {
		if strings.HasPrefix(raw, p) {
			return http.Get(raw)
		}
	}
	return nil, nil
}
