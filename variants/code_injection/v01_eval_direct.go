package main
import ("net/http"; "reflect")
func codeEval(w http.ResponseWriter, r *http.Request) {
    expr := r.FormValue("expr")
    _ = expr
    // reflect / plugin hooks often modeled as code paths
    _ = reflect.TypeOf(expr)
}
