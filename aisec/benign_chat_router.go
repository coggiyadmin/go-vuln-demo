// TN — benign chat router; routes by typed key, no untrusted text reaches a prompt.
package aisec

var handlers = map[string]func() string{
	"billing": func() string { return "Routing to billing." },
	"support": func() string { return "Routing to support." },
}

func Route(intent string) string {
	if h, ok := handlers[intent]; ok {
		return h()
	}
	return "unknown"
}
