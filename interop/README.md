# interop/ — in-language polyglot (Go)

In-language cross-language fixtures (host=Go embeds a guest DSL). Cross-binary cases live in
`coggiyadmin/interop-vuln-demo`. Plan:
`cogniumhq/sast-validation/research/cross-language-interop-plan.md` (IL-1).

Files use `//go:build ignore` or an isolated `package interop` so the main module still
builds. Each ships TP + `safe_*` + `benign_*`.

| Fixture | Boundary | CWE | Expected |
|---------|----------|-----|----------|
| `interop_htmltemplate.go` | Go → `template.HTML(user)` cast bypass | 79 | FN (cf #88c text/template) |
| `interop_sql_in_string.go` | Go → SQL DSL in string → `db.Query` | 89 | partial |
| `interop_shell_in_string.go` | Go → shell snippet → `exec.Command("sh","-c",…)` | 78 | partial |
