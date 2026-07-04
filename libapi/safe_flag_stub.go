# SUPERSEDED by real-dep replacement (INTAKE-56/57 N4) — kept for reference.
// TN — flag package spec, not shell (#167).
package libapi

import "flag"

func BuildFlagSet() *flag.FlagSet {
	return flag.NewFlagSet("demo", flag.ExitOnError)
}
