package main

import (
	"html/template"
	"net/http"
	"path/filepath"
)

// LayoutDir gives the location of the views
var LayoutDir string = "views/layouts"
var bootstrap *template.Template

func main() {
	var err error
	bootstrap, err = template.ParseFiles(layoutFiles()...)
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/", handler)
	http.ListenAndServe(":3000", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	bootstrap.ExecuteTemplate(w, "bootstrap", nil)
}

func layoutFiles() []string {
	files, err := filepath.Glob(LayoutDir + "/*.gohtml")
	if err != nil {
		panic(err)
	}
	return files
}
