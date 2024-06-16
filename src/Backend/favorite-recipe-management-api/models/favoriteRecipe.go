package models

import (
	"favorite-recipe-management.com/rest-api/db"
)

type FavoriteRecipe struct {
	ID       int64
	UserID   int64 `json:"userId" binding:"required"`
	RecipeID int64 `json:"recipeId" binding:"required"`
}

// Save creates a new FavoriteRecipe instance in the database
func (fav FavoriteRecipe) Save() error {
	query := "INSERT INTO favorite_recipes(user_id, recipe_id) VALUES ($1, $2) RETURNING id"

	err := db.DB.QueryRow(query, fav.UserID, fav.RecipeID).Scan(&fav.ID)
	if err != nil {
		return err
	}

	return nil
}

func GetAllFavoriteRecipesByUserId(userId int64) ([]FavoriteRecipe, error) {
	query := "SELECT * FROM favorite_recipes WHERE user_id = $1"
	rows, err := db.DB.Query(query, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var favoriteRecipes []FavoriteRecipe

	for rows.Next() {
		var fav FavoriteRecipe
		err := rows.Scan(&fav.ID, &fav.UserID, &fav.RecipeID)

		if err != nil {
			return nil, err
		}

		favoriteRecipes = append(favoriteRecipes, fav)
	}

	return favoriteRecipes, nil
}

func GetSingleFavoriteRecipe(favRecipeId int64) (*FavoriteRecipe, error) {
	query := "SELECT * FROM favorite_recipes WHERE id = $1"
	row := db.DB.QueryRow(query, favRecipeId)

	var favRecipe FavoriteRecipe
	err := row.Scan(&favRecipe.ID, &favRecipe.UserID, &favRecipe.RecipeID)

	if err != nil {
		return nil, err
	}

	return &favRecipe, nil
}

func (fav FavoriteRecipe) Delete() error {
	query := "DELETE FROM favorite_recipes WHERE id = $1"
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(fav.ID)
	return err
}

// Method to verify if user already favorited the recipe
func VerifyIfFavoriteRecipeExists(userId int64, recipeId int64) (bool, error) {
	query := "SELECT COUNT(*) FROM favorite_recipes WHERE user_id = $1 AND recipe_id = $2"
	var count int
	err := db.DB.QueryRow(query, userId, recipeId).Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}
