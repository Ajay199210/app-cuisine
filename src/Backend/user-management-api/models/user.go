package models

import (
	"errors"

	"usermanagement.com/rest-api/db"
	"usermanagement.com/rest-api/utils"
)

// CreateUserInput represents the input data required for creating a new user
type CreateUserInput struct {
	Email             string `json:"email" binding:"required"`
	Username          string `json:"username" binding:"required"`
	Password          string `json:"password" binding:"required"`
	PasswordConfirmed string `json:"passwordConfirmed" binding:"required"`
}

// LoginUserInput represents the input data required for user login
type LoginUserInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// ModifyUserInput represents the input data required for modifying user information
type ModifyUserInput struct {
	Username             string `json:"username,omitempty"`
	CurrentPassword      string `json:"currentPassword,omitempty"`
	NewPassword          string `json:"newPassword,omitempty"`
	NewPasswordConfirmed string `json:"newPasswordConfirmed,omitempty"`
}

// User represents the model for a user
type User struct {
	ID       int64
	Email    string
	Username string
	Password string
}

// Save creates a new user in the database
func (u *User) Save() error {
	// Hash the password before inserting it into the database
	hashedPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}

	// Define the SQL query to insert a new user
	query := "INSERT INTO users(email, username, password) VALUES ($1, $2, $3) RETURNING id"

	// Execute the query and retrieve the new user's ID
	err = db.DB.QueryRow(query, u.Email, u.Username, hashedPassword).Scan(&u.ID)
	if err != nil {
		return err
	}

	return nil
}

// ValidateCredentials checks if the provided username and password are valid for login
func (u User) ValidateCredentials(input LoginUserInput) error {
	query := "SELECT id, email, password FROM users WHERE username = $1"
	row := db.DB.QueryRow(query, input.Username)

	var retrievedEmail, retrievedPassword string
	err := row.Scan(&u.ID, &retrievedEmail, &retrievedPassword)

	if err != nil {
		return errors.New("credentials invalid")
	}

	passwordIsValid := utils.CheckPasswordHash(input.Password, retrievedPassword)

	if !passwordIsValid {
		return errors.New("credentials invalid")
	}

	u.Email = retrievedEmail
	u.Username = input.Username
	return nil
}

func GetAllUsers() ([]User, error) {
	query := "SELECT * FROM users"
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err //need to return nil and error because this function returns 2 values; so if error, nil will represent []Event and err for error
	}
	defer rows.Close()

	var users []User

	for rows.Next() { //this will loop through each row until there are none left
		var user User
		err := rows.Scan(&user.ID, &user.Email, &user.Username, &user.Password) //order of columns is important here

		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

func GetUserByID(id int64) (*User, error) {
	query := "SELECT * FROM users WHERE id = $1"
	row := db.DB.QueryRow(query, id) //QueryRow method will return the single row with the associated id

	var user User
	err := row.Scan(&user.ID, &user.Email, &user.Username, &user.Password)

	if err != nil {
		return nil, err
	}

	return &user, nil //returning a pointer instead of struct because if error is returned then pointer can be nil, but struct can't
}

func (u User) Update() error {
	query := `
	UPDATE users
	SET email = $1, username = $2, password = $3
	WHERE id = $4
	`
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	hashedPassword, err := utils.HashPassword(u.Password)

	if err != nil {
		return err
	}

	_, err = stmt.Exec(u.Email, u.Username, hashedPassword, u.ID) //make sure to respect order of query
	return err
}

func (u User) Delete() error {
	query := "DELETE FROM users WHERE id = $1"
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(u.ID)
	return err
}

// Method to check if a username already exists
func IsUsernameTaken(username string) (bool, error) {
	query := "SELECT COUNT(*) FROM users WHERE username = $1"
	var count int
	err := db.DB.QueryRow(query, username).Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}
