package context

import (
	"fmt"
	"net/http"
)

func Server(store Store) http.HandleFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		store.Cancel()
		fmt.Fprint(w, store.Fetch())
	}
}

// 
type Store interface {
	Fetch() string
	Cancel()
}
