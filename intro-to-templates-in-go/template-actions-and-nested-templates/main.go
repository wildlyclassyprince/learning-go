package main

import (
	"html/template"
	"net/http"
)

var testTemplate *template.Template

// Widget is simple testing struct
type Widget struct {
	Name  string
	Price int
}

// ViewData is simple testing struct
type ViewData struct {
	Name    string
	Widgets []Widget
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
	//err := testTemplate.Execute(w, nil)
	err := testTemplate.Execute(w, ViewData{
		Name: "John Smith",
		Widgets: []Widget{
			Widget{"Blue Widget", 12},
			Widget{"Red Widget", 15},
			Widget{"Green Widget", 17},
		},
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
