// v10 allowlist-set
package main
var safeIDs = map[int]bool{1: true, 2: true, 3: true}
func lookup(id int) map[string]int { if safeIDs[id] { return map[string]int{"id": id} }; return nil }
