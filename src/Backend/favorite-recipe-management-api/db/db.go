package db

import (
	"database/sql"
	"fmt"
	"log"
	"net/url"

	_ "github.com/lib/pq"
)

var DB *sql.DB

// function to initialize / create the database
func InitDB() {

	var err error
	// Encode the password
	encodedPassword := url.QueryEscape("F7)-wXwpSLkIkPFTKBY_BB5Zv~I>")

	// Construct the connection string
	connStr := fmt.Sprintf("postgres://postgres3:%s@db-rr-user-recipe-favorites.c14yie64wrh4.ca-central-1.rds.amazonaws.com:5432/postgres", encodedPassword)

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
	createFavoriteRecipesTable := `
	CREATE TABLE IF NOT EXISTS favorite_recipes (
		id SERIAL PRIMARY KEY,
		user_id TEXT NOT NULL,
		recipe_id TEXT NOT NULL
	)
	`

	_, err := DB.Exec(createFavoriteRecipesTable)

	if err != nil {
		panic("Could not create favorite_recipes table.")
	}
}
