package routes

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"user-recipe-management.com/rest-api/models"
)

func AddUserRecipe(context *gin.Context) {
	var userRecipe models.UserRecipe

	//binds the JSON data from the HTTP request to a UserRecipe struct
	err := context.ShouldBindJSON(&userRecipe)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "The following fields are required : userId, strMeal, strCategory, strInstructions, strIngredient1, strMeasure1"})
		return
	}

	// save the UserRecipe to the database
	err = userRecipe.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not save user recipe."})
		return
	}

	// Respond with a success message
	context.JSON(http.StatusCreated, gin.H{"message": "User recipe created successfully"})
}

func GetSingleUserRecipe(context *gin.Context) {
	recipeId, err := strconv.ParseInt(context.Param("recipeId"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse recipe id."})
		return
	}

	userRecipe, err := models.GetSingleUserRecipe(recipeId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch the recipe."})
		return
	}

	context.JSON(http.StatusOK, userRecipe)
}

func UpdateUserRecipe(context *gin.Context) {
	recipeId, err := strconv.ParseInt(context.Param("recipeId"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse recipe id."})
		return
	}

	userRecipe, err := models.GetSingleUserRecipe(recipeId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch the user's recipe."})
		return
	}

	var updatedUserRecipe models.UserRecipe
	err = context.ShouldBindJSON(&updatedUserRecipe)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "The following fields are required : userId, strMeal, strCategory, strInstructions, strIngredient1, strMeasure1"})
		return
	}

	updatedUserRecipe.IDMeal = userRecipe.IDMeal
	updatedUserRecipe.UserID = userRecipe.UserID
	err = updatedUserRecipe.Update()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not update the recipe."})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "User's recipe updated successfully!"})
}

func DeleteUserRecipe(context *gin.Context) {
	userRecipeID, err := strconv.ParseInt(context.Param("recipeId"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse recipe id."})
		return
	}

	userRecipe, err := models.GetSingleUserRecipe(userRecipeID)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch the user's recipe."})
		return
	}

	err = userRecipe.Delete()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not delete the user's recipe."})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "User's recipe deleted successfully!"})
}
