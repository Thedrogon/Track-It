package main

import (
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
	handler := Handlers.NewProblemHandler(repo)

	// Create router
	router := mux.NewRouter()

	// Serve static files
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Define routes
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/index.html")
	}).Methods("GET")
	router.HandleFunc("/problems", handler.CreateProblem).Methods("POST")
	router.HandleFunc("/problems", handler.GetAllProblems).Methods("GET")
	router.HandleFunc("/problems/{id}", handler.GetProblem).Methods("GET")
	router.HandleFunc("/problems/{id}", handler.UpdateProblem).Methods("PUT")
	router.HandleFunc("/problems/{id}", handler.DeleteProblem).Methods("DELETE")

	// Start server
	log.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
