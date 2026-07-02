// Combination #9 — COMMENT/STRING × NOSQL. Expect 0 findings.
package main

// user input never reaches sink — literal only
func commentNosql() {
	_ = "literal"
}
