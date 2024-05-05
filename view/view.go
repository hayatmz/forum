package view

import (
	"html/template"
	"net/http"
)

var tmpl *template.Template

func init() {
	tmpl = template.Must(template.ParseGlob("view/templates/*"))
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
}

func ExecTemplate(w http.ResponseWriter, tpl string, data any) {
	if tmpl.ExecuteTemplate(w, tpl, data) != nil {
		
	}
}