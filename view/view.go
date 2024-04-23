package view

import (
	"html/template"
)

func NewTemplate(file string) (*template.Template, error) {
	tmpl, errTmpl := template.ParseFiles("view/templates/"+file)
	if errTmpl != nil {
		return tmpl, errTmpl
	}
	return tmpl, nil
}