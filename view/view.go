package view

import (
	"html/template"
	"net/http"
)

var tmpl *template.Template

func init() {
	tmpl = template.Must(template.ParseGlob("view/templates/*"))
}

func NewTemplate(file string) (*template.Template, error) {
	tmpl, errTmpl := template.ParseFiles("view/templates/" + file)
	if errTmpl != nil {
		return tmpl, errTmpl
	}
	return tmpl, nil
}

func ExecTemplate(w http.ResponseWriter, tpl string, data any) error {
	if tmpl.ExecuteTemplate(w, tpl, data) != nil {
		tmpl.ExecuteTemplate(w, )
	}
}