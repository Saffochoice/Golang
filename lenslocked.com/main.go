package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html")
	if r.URL.Path == "/" {
		fmt.Fprint(w, "<h1>Welcome to main</h1>")
	} else if r.URL.Path == "/contact" {
		fmt.Fprint(w, "<h1>Welcome to contact</h1>")
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "<h1>PAGE NOT FOUND")
	}

}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", handlerFunc)
	http.ListenAndServe(":3010", r)
}
