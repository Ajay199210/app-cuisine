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
	connStr := fmt.Sprintf("postgres://postgres2:%s@db-rr-user-recipes.c14yie64wrh4.ca-central-1.rds.amazonaws.com:5432/postgres", encodedPassword)

	// Open connection to PostgreSQL
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	createTables()
}

// function to create necessary tables in the DB
// for this API, we only need a user recipe table for the moment
func createTables() {
	createUserRecipesTable := `
	CREATE TABLE IF NOT EXISTS user_recipes (
		id SERIAL PRIMARY KEY,
		user_id INTEGER NOT NULL,
		str_meal TEXT NOT NULL,
    	str_category TEXT NOT NULL,
    	str_area TEXT,
    	str_instructions TEXT NOT NULL,
    	str_meal_thumb TEXT,
    	str_ingredient1 TEXT NOT NULL,
    	str_ingredient2 TEXT,
    	str_ingredient3 TEXT,
    	str_ingredient4 TEXT,
    	str_ingredient5 TEXT,
    	str_ingredient6 TEXT,
    	str_ingredient7 TEXT,
    	str_ingredient8 TEXT,
    	str_ingredient9 TEXT,
    	str_ingredient10 TEXT,
    	str_ingredient11 TEXT,
    	str_ingredient12 TEXT,
    	str_ingredient13 TEXT,
    	str_ingredient14 TEXT,
    	str_ingredient15 TEXT,
    	str_ingredient16 TEXT,
    	str_ingredient17 TEXT,
    	str_ingredient18 TEXT,
    	str_ingredient19 TEXT,
    	str_ingredient20 TEXT,
    	str_measure1 TEXT NOT NULL,
    	str_measure2 TEXT,
    	str_measure3 TEXT,
    	str_measure4 TEXT,
    	str_measure5 TEXT,
    	str_measure6 TEXT,
    	str_measure7 TEXT,
    	str_measure8 TEXT,
    	str_measure9 TEXT,
    	str_measure10 TEXT,
    	str_measure11 TEXT,
    	str_measure12 TEXT,
    	str_measure13 TEXT,
    	str_measure14 TEXT,
    	str_measure15 TEXT,
    	str_measure16 TEXT,
    	str_measure17 TEXT,
    	str_measure18 TEXT,
    	str_measure19 TEXT,
    	str_measure20 TEXT
	)
	`
	_, err := DB.Exec(createUserRecipesTable)

	if err != nil {
		log.Fatalf("Could not create user_recipes table: %v", err) //I added this for debugging
		panic("Could not create user_recipes table.")
	}
}
