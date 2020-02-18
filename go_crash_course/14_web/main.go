package main

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world")
}

func about(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>About</h1>")
}

func main() {

	http.HandleFunc("/", index)
	http.HandleFunc("/about", about)

	fmt.Println("Serve Starting...")
	http.ListenAndServe(":3000", nil)
}
