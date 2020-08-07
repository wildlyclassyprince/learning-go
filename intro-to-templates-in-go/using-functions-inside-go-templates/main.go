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

func main() {
	var err error

	testTemplate, err = template.New("hello.gohtml").Funcs(template.FuncMap{
		"hasPermission": func(user User, feature string) bool {
			return false
		},
	}).ParseFiles("hello.gohtml")
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/", handler)
	http.ListenAndServe(":3000", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	user := User{
		ID:    1,
		Email: "jon@calhoun.io",
	}
	vd := ViewData{user}
	// We need to clone the template before setting a user-specific
	// FuncMap to avoid any potential race conditions.
	err := template.Must(testTemplate.Clone()).Funcs(template.FuncMap{
		"hasPermission": func(feature string) bool {
			if user.ID == 1 && feature == "feature-a" {
				return true
			}
			return false
		},
	}).Execute(w, vd)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
