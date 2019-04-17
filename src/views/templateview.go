package views

import (
	"html/template"
	"log"
	"middleware"
	"mux"
	"net/http"
)

var tmpl *template.Template

func init() {
	mux.GetMux().HandleFunc("/", middleware.WithMiddleware(templateview,
		middleware.Time(),
	))
}

func templateview(w http.ResponseWriter, r *http.Request) {
	tmpl := template.New("templateview")
	tmpl, err := tmpl.ParseFiles("tmpl/templateview.html")
	if err != nil {
		log.Println(err)
		log.Fatal("Template parsing error")
	}
	tmpl.ExecuteTemplate(w, "templateview", nil)
}
