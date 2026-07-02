// TN — text/template static merge (#156 velocity parity).
package libapi

import "text/template"

const velocityTpl = "<p>{{.Name}}</p>"

func RenderVelocity(userName string) (string, error) {
	tpl, err := template.New("greet").Parse(velocityTpl)
	if err != nil {
		return "", err
	}
	var buf stringsBuilder
	if err := tpl.Execute(&buf, map[string]string{"Name": userName}); err != nil {
		return "", err
	}
	return buf.String(), nil
}

type stringsBuilder struct{ b []byte }

func (s *stringsBuilder) Write(p []byte) (int, error) {
	s.b = append(s.b, p...)
	return len(p), nil
}

func (s *stringsBuilder) String() string { return string(s.b) }
