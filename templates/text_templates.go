package templates

import (
	"bytes"
	"log"
	"text/template"
)

type Person struct {
	Name       string
	Job string
}

const tmpl = `Name: {{.Name}}, Job: {{.Job}}`

func displayPerson(data Person) string {
	var t *template.Template
	t = template.New("mytemplate")
	t, err := t.Parse(tmpl)
	if err != nil {
		log.Fatal("Parse: ", err)
		return "ERROR"
	}
	var b bytes.Buffer
	err = t.Execute(&b, data)
	if err != nil {
		return "Error"
	}
	return b.String()
}

