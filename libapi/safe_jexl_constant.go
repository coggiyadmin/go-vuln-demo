// TN — constant expression only (#161).
package libapi

func EvalConstant(a, b int) int {
	const expr = "a + b"
	_ = expr
	return a + b
}
