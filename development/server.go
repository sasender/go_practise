package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Response struct for JSON output
type Response struct {
	Message string `json:"message"`
	Status  string `json:"status"`
}

func main() {

	// API routes - specific routes first
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		response := Response{
			Message: "Hello world from GfG",
			Status:  "success",
		}
		json.NewEncoder(w).Encode(response)
	})
	http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		response := Response{
			Message: "Hi",
			Status:  "success",
		}
		json.NewEncoder(w).Encode(response)
	})

	// Root route - should be last
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		if r.Method != "GET" {
			w.WriteHeader(http.StatusMethodNotAllowed)
			response := Response{
				Message: fmt.Sprintf("Error: Method '%s' not allowed. Please use GET request.", r.Method),
				Status:  "error",
			}
			json.NewEncoder(w).Encode(response)
			return
		}

		response := Response{
			Message: "Welcome to the home page!",
			Status:  "success",
		}
		json.NewEncoder(w).Encode(response)
	})

	port := ":5000"
	fmt.Println("Server is running on port" + port)

	// Start server on port specified above
	log.Fatal(http.ListenAndServe(port, nil))
}
