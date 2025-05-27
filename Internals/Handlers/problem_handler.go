package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/Thedrogon/Track-It/Internals/models"
	"github.com/Thedrogon/Track-It/Internals/repository"
	"github.com/gorilla/mux"
)

type ProblemHandler struct {
	repo *repository.ProblemRepository
}

func NewProblemHandler(repo *repository.ProblemRepository) *ProblemHandler {
	return &ProblemHandler{repo: repo}
}

func (h *ProblemHandler) Homepage(w http.ResponseWriter , r *http.Request){
	json.NewEncoder(w).Encode("Hello fellas")
}

func (h *ProblemHandler) CreateProblem(w http.ResponseWriter, r *http.Request) {
	var problem models.Problem
	if err := json.NewDecoder(r.Body).Decode(&problem); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.repo.Create(&problem); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(problem)
}

func (h *ProblemHandler) GetProblem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	problem, err := h.repo.GetByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(problem)
}

func (h *ProblemHandler) GetAllProblems(w http.ResponseWriter, r *http.Request) {
	problems, err := h.repo.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(problems)
}

func (h *ProblemHandler) UpdateProblem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var problem models.Problem
	if err := json.NewDecoder(r.Body).Decode(&problem); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	problem.ID = id
	if err := h.repo.Update(&problem); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(problem)
}

func (h *ProblemHandler) DeleteProblem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	if err := h.repo.Delete(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *ProblemHandler) GetProblemsByTags(w http.ResponseWriter, r *http.Request) {
	// Get tags from query parameter
	tagsParam := r.URL.Query().Get("tags")
	if tagsParam == "" {
		http.Error(w, "tags parameter is required", http.StatusBadRequest)
		return
	}

	// Split tags by comma and trim whitespace
	tags := strings.Split(tagsParam, ",")
	for i := range tags {
		tags[i] = strings.TrimSpace(tags[i])
	}

	problems, err := h.repo.GetByTags(tags)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(problems)
}
