// FP-target (N4 real dep #167) — flag.FlagSet spec, not shell.
package libapi

import "flag"

func BuildFlagSet() *flag.FlagSet {
	return flag.NewFlagSet("demo", flag.ContinueOnError)
}
