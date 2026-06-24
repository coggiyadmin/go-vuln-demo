// FP-target (upstream cognium-dev#162/#140) — DEEP reflective-load flow. A dev/CLI plugin loader
// dynamically loads a Go plugin (.so) by path and resolves a symbol via reflection-style Lookup,
// then invokes it. `soPath`/`symbol` are operator-controlled tool arguments, not an HTTP source,
// so this must NOT be flagged code_injection under an entry-point gate — even though a real
// dynamic-load + dynamic-dispatch sink is present.
package profile

import "plugin"

// LoadAndRun opens a plugin shipped/selected by the operator and calls its Run symbol.
func LoadAndRun(soPath, symbol string) error {
	p, err := plugin.Open(soPath) // operator-controlled path, not request data
	if err != nil {
		return err
	}
	sym, err := p.Lookup(symbol) // dynamic symbol resolution (reflection-style)
	if err != nil {
		return err
	}
	if run, ok := sym.(func() error); ok {
		return run()
	}
	return nil
}
