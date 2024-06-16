package routes

import (
	"net/http"
	"strconv"
	"unicode"

	"github.com/gin-gonic/gin"
	"usermanagement.com/rest-api/models"
	"usermanagement.com/rest-api/utils"
)

func Signup(context *gin.Context) {
	var user models.CreateUserInput // Use CreateUserInput for signup

	// Bind the request JSON data to the CreateUserInput struct
	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
		return
	}

	//Check that username isn't already taken
	taken, err := models.IsUsernameTaken(user.Username)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Error checking username availability."})
		return
	}

	//If username already exists in database, then bad request + message is returned
	if taken {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Username already taken. Please choose another"})
		return
	}

	// Check if password and password confirmation match
	if user.Password != user.PasswordConfirmed {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Passwords do not match."})
		return
	}

	// Check password complexity
	if !passwordComplexityCheck(user.Password) {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Password must be at least 8 characters long and include uppercase, lowercase letters, numbers, and special characters."})
		return
	}

	// Create a new User instance from the input data
	newUser := models.User{
		Email:    user.Email,
		Username: user.Username,
		Password: user.Password,
	}

	// Save the user to the database
	err = newUser.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not save user."})
		return
	}

	// Respond with a success message if user is saved to database
	context.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}

func Login(context *gin.Context) {
	var user models.LoginUserInput

	// Bind the request JSON data to the LoginUserInput struct
	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
		return
	}

	// Create a new User instance to use the ValidateCredentials method
	loginUser := models.User{
		Username: user.Username,
		Password: user.Password,
	}

	// Validate the user's credentials
	err = loginUser.ValidateCredentials(user)

	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Could not authenticate user."})
		return
	}

	// Generate a token for the authenticated user
	token, err := utils.GenerateToken(loginUser.Username, loginUser.ID)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not authenticate user."})
		return
	}

	//Set cookie for user session
	context.SetSameSite(http.SameSiteLaxMode)
	context.SetCookie("Authorization", token, 3600*24*30, "", "localhost", false, true)

	// Respond with a success message and the generated token
	context.JSON(http.StatusOK, gin.H{"message": "Login successful!", "token": token})
}

func Logout(context *gin.Context) {
	context.SetSameSite(http.SameSiteLaxMode)
	context.SetCookie("Authorization", "", -1, "", "localhost", false, true)
	context.JSON(http.StatusOK, gin.H{"message": "Logout successful!"})
}

func GetUsers(context *gin.Context) {
	users, err := models.GetAllUsers()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch users. Try again later."})
		return
	}

	// Construct a response containing only necessary user data (no password)
	var usersResponse []gin.H
	for _, user := range users {
		userResponse := gin.H{
			"id":       user.ID,
			"email":    user.Email,
			"username": user.Username,
		}
		usersResponse = append(usersResponse, userResponse)
	}

	context.JSON(http.StatusOK, usersResponse)
}

func GetUser(context *gin.Context) {
	userId, err := strconv.ParseInt(context.Param("id"), 10, 64) //.Param gets path paramter as a string so that's why we parse it
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse user id."})
		return
	}

	user, err := models.GetUserByID(userId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch user."})
		return
	}

	// Return a response containing only necessary user data (no password)
	context.JSON(http.StatusOK, gin.H{
		"id":       user.ID,
		"email":    user.Email,
		"username": user.Username,
	})
}

// Update function allows users to update password and username
func UpdateUser(context *gin.Context) {
	userId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse user id."})
		return
	}

	user, err := models.GetUserByID(userId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch the user."})
		return
	}

	var input models.ModifyUserInput
	err = context.ShouldBindJSON(&input)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
		return
	}

	// Check if the username is being changed and if it's unique
	if input.Username != "" && input.Username != user.Username {
		taken, err := models.IsUsernameTaken(input.Username)
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"message": "Error checking username availability."})
			return
		}

		if taken {
			context.JSON(http.StatusBadRequest, gin.H{"message": "Username already taken. Please choose another"})
			return
		}
		user.Username = input.Username
	}

	// Check if the password is being changed
	if input.NewPassword != "" {
		// Verify the current password
		//User can only change password if the current one is entered first and is correct
		currentPasswordVerified := utils.CheckPasswordHash(input.CurrentPassword, user.Password)

		if !currentPasswordVerified {
			context.JSON(http.StatusBadRequest, gin.H{"message": "Current password is incorrect."})
			return
		}

		// Check if the new password matches the confirmation and meets complexity requirements
		if input.NewPassword != input.NewPasswordConfirmed {
			context.JSON(http.StatusBadRequest, gin.H{"message": "New password and confirmation do not match."})
			return
		}
		if !passwordComplexityCheck(input.NewPassword) {
			context.JSON(http.StatusBadRequest, gin.H{"message": "Password must be at least 8 characters long and include uppercase, lowercase letters, numbers, and special characters."})
			return
		}

		user.Password = input.NewPassword
	}

	// Update the user in the database
	err = user.Update()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not update user."})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "User updated successfully!"})
}

func DeleteUser(context *gin.Context) {
	userId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse user id."})
		return
	}

	user, err := models.GetUserByID(userId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch the user."})
		return
	}

	err = user.Delete()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not delete the user."})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "User deleted successfully!"})
}

// Function to check password complexity
// Had issues with regular expression so read that this approach is recommended
func passwordComplexityCheck(password string) bool {
	// letters := 0
	chars := 0
	var number, upper, special, eightOrMore bool = false, false, false, false

	for _, c := range password {
		switch {
		case unicode.IsNumber(c):
			number = true
			chars++
		case unicode.IsUpper(c):
			upper = true
			chars++
		case unicode.IsPunct(c) || unicode.IsSymbol(c):
			special = true
			chars++
		case unicode.IsLetter(c) || c == ' ':
			chars++
		default:
			return false
		}
	}

	eightOrMore = chars >= 8

	if number && upper && special && eightOrMore {
		return true
	}

	return false
}
