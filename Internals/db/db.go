package database

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "./problems.db")
	if err != nil {
		log.Fatal(err)
	}

	// Create problems table if it doesn't exist
	createTable := `
	CREATE TABLE IF NOT EXISTS problems (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		problem_id INTEGER NOT NULL,
		title TEXT NOT NULL,
		tags TEXT
	);`

	_, err = DB.Exec(createTable)
	if err != nil {
		log.Fatal(err)
	}
}
