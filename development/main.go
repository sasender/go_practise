package main

import (
	"fmt"
	"log"
	"net/http"
)

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "ParseForm() error: "+err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "POST request successful\n")

	username := r.FormValue("username")
	address := r.FormValue("address")
	state := r.FormValue("state")

	fmt.Fprintf(w, "Username = %s\n", username)
	fmt.Fprintf(w, "Address = %s\n", address)
	fmt.Fprintf(w, "State = %s\n", state)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method not supported", http.StatusMethodNotAllowed) // Fixed status
		return
	}
	fmt.Fprintf(w, "Hello, World!")
}

func main() {
	http.HandleFunc("/form", formHandler) // Specific first
	http.HandleFunc("/hello", helloHandler)
	http.Handle("/", http.FileServer(http.Dir("./static"))) // Catch-all last

	fmt.Printf("Starting Server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
