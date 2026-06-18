package sstixf

import (
	"os"
	"text/template"
)

func Run(name string) {
	t, _ := template.New("t").Parse("<p>Hello " + name + "</p>") // SINK CWE-1336
	t.Execute(os.Stdout, nil)
}
