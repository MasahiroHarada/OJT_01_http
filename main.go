package main

import (
	"net/http"
	"html/template"
	"log"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/reset", reset)

	// static contents
	http.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("./static"))))

	http.ListenAndServe(":9876", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		// TODO session
		err := tpl.ExecuteTemplate(w, "index.html", nil)
		handleError(w, err)
	case http.MethodPost:
		req.ParseForm()
		var name = req.PostFormValue("username")
		// TODO session
		err := tpl.ExecuteTemplate(w, "index.html", struct {
			Name string
		}{
			Name: name,
		})
		handleError(w, err)
	}
}

func reset(w http.ResponseWriter, req *http.Request) {
	// TODO reset session

	// TODO redirect to /
}

func handleError(w http.ResponseWriter, err error) {
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Fatalln(err.Error())
	}
}
