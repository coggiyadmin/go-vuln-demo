// TN — html/template static parse; user data at Execute (#156).
package libapi

import "html/template"

const tplSource = "<p>{{.Name}}</p>"

func RenderTemplate(userName string) (string, error) {
	tpl, err := template.New("greet").Parse(tplSource)
	if err != nil {
		return "", err
	}
	var out stringsBuilder
	err = tpl.Execute(&out, map[string]string{"Name": userName})
	return out.String(), err
}

type stringsBuilder struct{ b []byte }

func (s *stringsBuilder) Write(p []byte) (int, error) {
	s.b = append(s.b, p...)
	return len(p), nil
}

func (s *stringsBuilder) String() string { return string(s.b) }
