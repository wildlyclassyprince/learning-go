package main

import (
	"html/template"
	"net/http"
)

var testTemplate *template.Template

// ViewData struct for testing data view
type ViewData struct {
	User User
}

// User struct just for testing
type User struct {
	ID    int
	Email string
}

// HasPermission determines
func (u User) HasPermission(feature string) bool {
	if feature == "feature-a" {
		return true
	} else {
		return false
	}
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

	vd := ViewData{
		User: User{1, "jon@calhoun.io"},
	}
	err := testTemplate.Execute(w, vd)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
