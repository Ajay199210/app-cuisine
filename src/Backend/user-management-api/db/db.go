package db

import (
	"database/sql"
	"fmt"
	"log"
	"net/url"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() {
	var err error

	// Encode the password
	encodedPassword := url.QueryEscape("F7)-wXwpSLkIkPFTKBY_BB5Zv~I>")

	// Construct the connection string
	connStr := fmt.Sprintf("postgres://postgres:%s@db-user-management-recette-remix.c14yie64wrh4.ca-central-1.rds.amazonaws.com:5432/", encodedPassword)

	// Open connection to PostgreSQL
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	createTables()
}

func createTables() {
	createUsersTable := `
	CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		email TEXT NOT NULL,
		username TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL
	)
	`

	_, err := DB.Exec(createUsersTable)

	if err != nil {
		panic("Could not create users table.")
	}
}
