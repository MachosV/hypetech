package views

import (
	"html/template"
	"log"
	"mux"
	"net/http"
)

var tmpl *template.Template

func init() {
	var err error
	mux.GetMux().HandleFunc("/", getdata)
	tmpl = template.New("getdata")
	tmpl, err = tmpl.ParseFiles("tmpl/getdata.html")
	if err != nil {
		log.Println(err)
		log.Fatal("Template parsing error")
	}
}

func getdata(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "getdata", nil)
}
