// TN — flag package spec, not shell (#167).
package libapi

import "flag"

func BuildFlagSet() *flag.FlagSet {
	return flag.NewFlagSet("demo", flag.ExitOnError)
}
