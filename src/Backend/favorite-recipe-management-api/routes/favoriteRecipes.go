package routes

import (
	"net/http"
	"strconv"

	"favorite-recipe-management.com/rest-api/models"
	"github.com/gin-gonic/gin"
)

func AddFavoriteRecipe(context *gin.Context) {
	var favRecipe models.FavoriteRecipe

	//binds the JSON data from the HTTP request to a FavoriteRecipe struct
	err := context.ShouldBindJSON(&favRecipe)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "The following fields are required : userId, recipeId"})
		return
	}

	alreadyExists, err := models.VerifyIfFavoriteRecipeExists(favRecipe.UserID, favRecipe.RecipeID)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Error verifying if favorite recipe already exists."})
		return
	}

	//If favorite recipe for user already exists in database, then bad request + message is returned
	if alreadyExists {
		context.JSON(http.StatusBadRequest, gin.H{"message": "This recipe has already been added to user's favorites"})
		return
	}

	// Save the FavoriteRecipe to the database
	err = favRecipe.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not save the favorited recipe."})
		return
	}

	// Respond with a success message
	context.JSON(http.StatusCreated, gin.H{"message": "Favorite recipe created successfully"})
}

func GetAllFavoriteRecipesByUserId(context *gin.Context) {
	userId, err := strconv.ParseInt(context.Param("userId"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse user id."})
		return
	}

	favRecipes, err := models.GetAllFavoriteRecipesByUserId(userId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch user's favorite recipes. Try again later."})
		return
	}

	context.JSON(http.StatusOK, favRecipes)
}

func DeleteFavoriteRecipe(context *gin.Context) {
	favRecipeID, err := strconv.ParseInt(context.Param("favRecipeId"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse recipe id."})
		return
	}

	favRecipe, err := models.GetSingleFavoriteRecipe(favRecipeID)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch the user's favorite recipe."})
		return
	}

	err = favRecipe.Delete()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not delete the user's favorited recipe."})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "User's favorite recipe deleted successfully!"})
}
