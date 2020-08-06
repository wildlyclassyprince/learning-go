package main

import (
	"html/template"
	"net/http"
)

var testTemplate *template.Template

// User struct for testing
type User struct {
	Admin bool
}

// ViewData struct for testing data view
type ViewData struct {
	*User
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

	vd := ViewData{&User{false}}
	err := testTemplate.Execute(w, vd)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
