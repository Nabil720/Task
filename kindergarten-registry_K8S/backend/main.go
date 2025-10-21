package main

import (
	"backend/handlers"
	"log"
	"net/http"
)

// CORS middleware function - FIXED VERSION
func enableCors(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Simple solution - allow all origins in development
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// Handle preflight OPTIONS request
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		// Call the next handler
		next(w, r)
	}
}

// Health check endpoint
func healthz(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

// Readiness check endpoint
func readiness(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Ready"))
}

func main() {
	// Apply CORS middleware to each route
	http.HandleFunc("/add-student", enableCors(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			handlers.AddStudent(w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	}))

	http.HandleFunc("/students", enableCors(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			handlers.GetStudents(w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	}))

	http.HandleFunc("/delete-student", enableCors(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodDelete {
			handlers.DeleteStudent(w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	}))

	http.HandleFunc("/update-student", enableCors(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPut {
			handlers.UpdateStudent(w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	}))

	// Health and Readiness endpoints
	http.HandleFunc("/healthz", healthz)
	http.HandleFunc("/readiness", readiness)

	// Start the server
	log.Println("Server running on port 5000")
	log.Fatal(http.ListenAndServe(":5000", nil))
}