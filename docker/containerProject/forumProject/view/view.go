package view

import (
	"html/template"
	"net/http"
)

var tmpl *template.Template

// parse all templates
func init() {
	tmpl = template.Must(template.ParseGlob("view/templates/*"))
}

// execute the template with the data, if the template contain an error execute the error template or a http error
//
// execute with the header connected or not connected in relation to the tpl argument
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

// add the main data and the tmpl string argument to the data to serve at the template
func (allDATA *DataTMPL) loadAllDATA(w http.ResponseWriter, data any) {
	allDATA.Tmpl = w.Header().Get("connected")
	allDATA.Data = data
}