package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Thedrogon/Track-It/Internals/db"
	"github.com/Thedrogon/Track-It/Internals/handlers"
	"github.com/Thedrogon/Track-It/Internals/repository"
	"github.com/gorilla/mux"
)

func main() {
	// Initialize database
	db.InitDB()

	// Initialize repository and handler
	repo := repository.NewProblemRepository()
	handler := handlers.NewProblemHandler(repo)

	// Create router
	router := mux.NewRouter()

	// API Documentation endpoint
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		apiDocs := map[string]string{
			"API Endpoints":                     "Available endpoints for the Problem Tracker API:",
			"GET /problems":                     "Get all problems",
			"GET /problems/tags?tags=tag1,tag2": "Get problems by tags (comma-separated)",
			"POST /problems":                    "Create a new problem (requires JSON body with problem_id, title, and tags)",
			"GET /problems/{id}":                "Get a specific problem by ID",
			"PUT /problems/{id}":                "Update a problem (requires JSON body with problem_id, title, and tags)",
			"DELETE /problems/{id}":             "Delete a problem",
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(apiDocs)
	}).Methods("GET")

	// API routes
	router.HandleFunc("/problems", handler.CreateProblem).Methods("POST")
	router.HandleFunc("/problems", handler.GetAllProblems).Methods("GET")
	router.HandleFunc("/problems/tags", handler.GetProblemsByTags).Methods("GET")
	router.HandleFunc("/problems/{id}", handler.GetProblem).Methods("GET")
	router.HandleFunc("/problems/{id}", handler.UpdateProblem).Methods("PUT")
	router.HandleFunc("/problems/{id}", handler.DeleteProblem).Methods("DELETE")

	// Start server
	log.Println("Server starting on :8080")
	log.Println("API Documentation available at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
