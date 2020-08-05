package main

import (
	"html/template"
	"net/http"
)

var testTemplate *template.Template

// ViewData is simple testing struct
type ViewData struct {
	Name string
}

func main() {
	var err error
	testTemplate, err = template.ParseFiles("hello.gohtml")
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/", handler)
	http.ListenAndServe(":3000", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	//vd := ViewData{"John Smith"}
	//err := testTemplate.Execute(w, vd)
	err := testTemplate.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
