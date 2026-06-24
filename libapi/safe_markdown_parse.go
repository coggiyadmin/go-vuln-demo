// FP-target (upstream cognium-dev#155, go) — a Markdown parser's Parse() produces an AST, it does
// NOT execute code. Must not be flagged code_injection.
package libapi

// MarkdownParser is a goldmark-style parser; Parse builds an AST node tree.
type MarkdownParser interface{ Parse(src []byte) Node }
type Node interface{ Kind() string }

// Render parses untrusted markdown into an AST. Parsing is not code execution.
func Render(p MarkdownParser, markdown string) Node {
	return p.Parse([]byte(markdown)) // AST build — NOT code execution
}
