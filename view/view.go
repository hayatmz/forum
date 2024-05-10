package view

import (
	"html/template"
	"net/http"
)

var tmpl *template.Template

func init() {
	// http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.Handle("/view/static/", http.StripPrefix("/view/static/", http.FileServer(http.Dir("view/static"))))
	tmpl = template.Must(template.ParseGlob("view/templates/*"))
	// http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("view/static"))))
}

func ExecTemplate(w http.ResponseWriter, tpl string, data any) {
	err := tmpl.ExecuteTemplate(w, tpl, data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		err := tmpl.ExecuteTemplate(w, "error", http.StatusInternalServerError)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}