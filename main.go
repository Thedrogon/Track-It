package main

import (
	"log"
	"net/http"

	handlers "github.com/Thedrogon/Track-It/Internals/Handlers"
	"github.com/Thedrogon/Track-It/Internals/repository"
	"github.com/gorilla/mux"
	"github.com/Thedrogon/Track-It/Internals/db"
)

func main() {
	// Initialize database
	database.InitDB()

	// Initialize repository and handler
	repo := repository.NewProblemRepository()
	handler := handlers.NewProblemHandler(repo)

	// Create router
	router := mux.NewRouter()

	// Define routes
	router.HandleFunc("/problems", handler.CreateProblem).Methods("POST")
	router.HandleFunc("/problems", handler.GetAllProblems).Methods("GET")
	router.HandleFunc("/problems/{id}", handler.GetProblem).Methods("GET")
	router.HandleFunc("/problems/{id}", handler.UpdateProblem).Methods("PUT")
	router.HandleFunc("/problems/{id}", handler.DeleteProblem).Methods("DELETE")

	// Start server
	log.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
