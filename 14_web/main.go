package main

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home page")
}

func about(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "About page")
}

func main() {
  
	http.HandleFunc("/", index)
	http.HandleFunc("/about", about)

	fmt.Println("Server running")

	http.ListenAndServe(":3000", nil)

}