package routes

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"user-recipe-management.com/rest-api/models"
)

func GetAllUserRecipesByUserID(context *gin.Context) {
	userId, err := strconv.ParseInt(context.Param("userId"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse user id."})
		return
	}

	userRecipes, err := models.GetAllUserRecipesByUserID(userId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch user's recipes. Try again later."})
		return
	}

	context.JSON(http.StatusOK, userRecipes)
}
