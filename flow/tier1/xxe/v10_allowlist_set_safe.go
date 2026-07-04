// v10 allowlist-set
package main
import "encoding/xml"
func parse(raw []byte) error { return xml.Unmarshal(raw, &struct{}{}) }
