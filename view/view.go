package view

import (
	"fmt"
	"html/template"
)

func NewTemplate() *template.Template {
	tmpl, errTmpl := template.ParseFiles("view/templates/index.html")
	if errTmpl != nil {
		fmt.Println("err", errTmpl)
	}
	return tmpl
}