// v09 guard-in-variable — host allowlist in package map.
package main

import ("net/http"; "net/url")

var allowedHosts = map[string]bool{"api.internal.example.com": true}

func fetch(raw string) (*http.Response, error) {
	u, err := url.Parse(raw)
	if err != nil {
		return nil, err
	}
	if allowedHosts[u.Hostname()] {
		return http.Get(raw)
	}
	return nil, nil
}
