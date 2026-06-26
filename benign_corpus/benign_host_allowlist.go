package main

import ("net/http"; "net/url")

func fetch(raw string) (*http.Response, error) {
	u, err := url.Parse(raw)
	if err != nil || u.Hostname() != "api.internal.example.com" {
		return nil, err
	}
	return http.Get(raw)
}
