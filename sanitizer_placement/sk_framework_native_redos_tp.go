package main
import ("html"; "net/http"; "regexp")
func skFrameworkNativeRedos(w http.ResponseWriter, r *http.Request) {
    pat := html.EscapeString(r.URL.Query().Get('p'))
    re, _ := regexp.Compile(pat)
    re.MatchString(r.URL.Query().Get('t'))
}
