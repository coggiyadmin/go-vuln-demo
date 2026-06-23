// Provenance/Misinfo (OWASP LLM09) — hallucinated dependency (slopsquatting surface).
package aisec

import retry "example.com/hallucinated/httpretry" // SINK (LLM09): fabricated module path

func FetchURL(url string) (string, error) { return retry.Get(url) }
