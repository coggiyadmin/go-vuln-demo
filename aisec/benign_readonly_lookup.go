// TN — benign read-only lookup; fixed catalog by exact key.
package aisec

import "strings"

var catalog = map[string]string{"USD": "US Dollar", "EUR": "Euro", "JPY": "Japanese Yen"}

func CurrencyName(code string) string {
	if name, ok := catalog[strings.ToUpper(code)]; ok {
		return name
	}
	return "unknown"
}
