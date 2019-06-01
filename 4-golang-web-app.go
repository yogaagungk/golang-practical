package main

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Go is neat")
}

func about(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Go is an open source programming language that makes it easy to build simple, reliable, and efficient software.")
}

// GolangWebApp : serve web app
func GolangWebApp() {

	http.HandleFunc("/", index)
	http.HandleFunc("/about", about)
	http.ListenAndServe(":8080", nil)
}
