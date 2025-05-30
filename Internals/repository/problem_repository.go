package repository

import (
	"database/sql"
	"encoding/json"
	"errors"

	"github.com/Thedrogon/Track-It/Internals/db"
	"github.com/Thedrogon/Track-It/Internals/models"
)

type ProblemRepository struct {
	db *sql.DB
}

func NewProblemRepository() *ProblemRepository {
	return &ProblemRepository{
		db: db.DB,
	}
}

func (r *ProblemRepository) Create(problem *models.Problem) error {
	tagsJSON, err := json.Marshal(problem.Tags)
	if err != nil {
		return err
	}

	query := `INSERT INTO problems (problem_id, title, tags) VALUES (?, ?, ?)`
	result, err := r.db.Exec(query, problem.Problem_ID, problem.Title, string(tagsJSON))
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	problem.ID = int(id)
	return nil
}

func (r *ProblemRepository) GetByID(id int) (*models.Problem, error) {
	problem := &models.Problem{}
	var tagsJSON string

	query := `SELECT id, problem_id, title, tags FROM problems WHERE id = ?`
	err := r.db.QueryRow(query, id).Scan(&problem.ID, &problem.Problem_ID, &problem.Title, &tagsJSON)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("problem not found")
		}
		return nil, err
	}

	err = json.Unmarshal([]byte(tagsJSON), &problem.Tags)
	if err != nil {
		return nil, err
	}

	return problem, nil
}

func (r *ProblemRepository) GetAll() ([]*models.Problem, error) {
	query := `SELECT id, problem_id, title, tags FROM problems`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var problems []*models.Problem
	for rows.Next() {
		problem := &models.Problem{}
		var tagsJSON string
		err := rows.Scan(&problem.ID, &problem.Problem_ID, &problem.Title, &tagsJSON)
		if err != nil {
			return nil, err
		}

		err = json.Unmarshal([]byte(tagsJSON), &problem.Tags)
		if err != nil {
			return nil, err
		}

		problems = append(problems, problem)
	}

	return problems, nil
}

func (r *ProblemRepository) Update(problem *models.Problem) error {
	tagsJSON, err := json.Marshal(problem.Tags)
	if err != nil {
		return err
	}

	query := `UPDATE problems SET problem_id = ?, title = ?, tags = ? WHERE id = ?`
	_, err = r.db.Exec(query, problem.Problem_ID, problem.Title, string(tagsJSON), problem.ID)
	return err
}

func (r *ProblemRepository) Delete(id int) error {
	query := `DELETE FROM problems WHERE id = ?`
	_, err := r.db.Exec(query, id)
	return err
}

func (r *ProblemRepository) GetByTags(tags []string) ([]*models.Problem, error) {
	query := `SELECT id, problem_id, title, tags FROM problems`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var problems []*models.Problem
	for rows.Next() {
		problem := &models.Problem{
			Tags: []string{},
		}
		var tagsJSON string
		err := rows.Scan(&problem.ID, &problem.Problem_ID, &problem.Title, &tagsJSON)
		if err != nil {
			return nil, err
		}

		err = json.Unmarshal([]byte(tagsJSON), &problem.Tags)
		if err != nil {
			return nil, err
		}

		// Check if problem has any of the requested tags
		hasTag := false
		for _, tag := range tags {
			for _, problemTag := range problem.Tags {
				if tag == problemTag {
					hasTag = true
					break
				}
			}
			if hasTag {
				break
			}
		}

		if hasTag {
			problems = append(problems, problem)
		}
	}

	return problems, nil
}
