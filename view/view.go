package view

import (
	"html/template"
	"net/http"
)

var tmpl *template.Template

func init() {
	tmpl = template.Must(template.ParseGlob("view/templates/*"))
}

func ExecTemplate(w http.ResponseWriter, tpl, msg string, data any) {
	var allDATA DataTMPL
	allDATA.loadAllDATA(w, data)
	allDATA.Msg = msg

	err := tmpl.ExecuteTemplate(w, tpl, allDATA)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		err := tmpl.ExecuteTemplate(w, "error.html", http.StatusInternalServerError)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

func (allDATA *DataTMPL) loadAllDATA(w http.ResponseWriter, data any) {
	allDATA.Tmpl = w.Header().Get("connected")
	allDATA.Data = data
}