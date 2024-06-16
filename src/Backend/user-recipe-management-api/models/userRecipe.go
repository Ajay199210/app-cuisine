package models

import "user-recipe-management.com/rest-api/db"

//UserRecipe represents the model of a recipe created by a user
//Replicated the same recipe object that's returned by The MealDB
//However, I decided to remove the following MealDB fields : strDrinkAlternate, strTags, strYoutube, strSource, strImageSource, strCreaticveCommonsConfirmed, dateModified
type UserRecipe struct {
	IDMeal          int64  `json:"idMeal"`
	UserID          int64  `json:"userId" binding:"required"`
	StrMeal         string `json:"strMeal" binding:"required"`
	StrCategory     string `json:"strCategory" binding:"required"`
	StrArea         string `json:"strArea"`
	StrInstructions string `json:"strInstructions" binding:"required"`
	StrMealThumb    string `json:"strMealThumb"`
	StrIngredient1  string `json:"strIngredient1" binding:"required"`
	StrIngredient2  string `json:"strIngredient2"`
	StrIngredient3  string `json:"strIngredient3"`
	StrIngredient4  string `json:"strIngredient4"`
	StrIngredient5  string `json:"strIngredient5"`
	StrIngredient6  string `json:"strIngredient6"`
	StrIngredient7  string `json:"strIngredient7"`
	StrIngredient8  string `json:"strIngredient8"`
	StrIngredient9  string `json:"strIngredient9"`
	StrIngredient10 string `json:"strIngredient10"`
	StrIngredient11 string `json:"strIngredient11"`
	StrIngredient12 string `json:"strIngredient12"`
	StrIngredient13 string `json:"strIngredient13"`
	StrIngredient14 string `json:"strIngredient14"`
	StrIngredient15 string `json:"strIngredient15"`
	StrIngredient16 string `json:"strIngredient16"`
	StrIngredient17 string `json:"strIngredient17"`
	StrIngredient18 string `json:"strIngredient18"`
	StrIngredient19 string `json:"strIngredient19"`
	StrIngredient20 string `json:"strIngredient20"`
	StrMeasure1     string `json:"strMeasure1" binding:"required"`
	StrMeasure2     string `json:"strMeasure2"`
	StrMeasure3     string `json:"strMeasure3"`
	StrMeasure4     string `json:"strMeasure4"`
	StrMeasure5     string `json:"strMeasure5"`
	StrMeasure6     string `json:"strMeasure6"`
	StrMeasure7     string `json:"strMeasure7"`
	StrMeasure8     string `json:"strMeasure8"`
	StrMeasure9     string `json:"strMeasure9"`
	StrMeasure10    string `json:"strMeasure10"`
	StrMeasure11    string `json:"strMeasure11"`
	StrMeasure12    string `json:"strMeasure12"`
	StrMeasure13    string `json:"strMeasure13"`
	StrMeasure14    string `json:"strMeasure14"`
	StrMeasure15    string `json:"strMeasure15"`
	StrMeasure16    string `json:"strMeasure16"`
	StrMeasure17    string `json:"strMeasure17"`
	StrMeasure18    string `json:"strMeasure18"`
	StrMeasure19    string `json:"strMeasure19"`
	StrMeasure20    string `json:"strMeasure20"`
}

// Save creates a new UserRecipe in the database
func (u *UserRecipe) Save() error {
	query := ` 
	INSERT INTO user_recipes (
		user_id,
		str_meal,
		str_category,
		str_area,
		str_instructions,
		str_meal_thumb,
		str_ingredient1,
		str_ingredient2,
		str_ingredient3,
		str_ingredient4,
		str_ingredient5,
		str_ingredient6,
		str_ingredient7,
		str_ingredient8,
		str_ingredient9,
		str_ingredient10,
		str_ingredient11,
		str_ingredient12,
		str_ingredient13,
		str_ingredient14,
		str_ingredient15,
		str_ingredient16,
		str_ingredient17,
		str_ingredient18,
		str_ingredient19,
		str_ingredient20,
		str_measure1,
		str_measure2,
		str_measure3,
		str_measure4,
		str_measure5,
		str_measure6,
		str_measure7,
		str_measure8,
		str_measure9,
		str_measure10,
		str_measure11,
		str_measure12,
		str_measure13,
		str_measure14,
		str_measure15,
		str_measure16,
		str_measure17,
		str_measure18,
		str_measure19,
		str_measure20 ) 
	VALUES (
		$1,$2,$3,$4,$5,$6,$7,$8,$9,$10,
		$11,$12,$13,$14,$15,$16,$17,$18,$19,$20,
		$21,$22,$23,$24,$25,$26,$27,$28,$29,$30,
		$31,$32,$33,$34,$35,$36,$37,$38,$39,$40,
		$41,$42,$43,$44,$45,$46
		)
	RETURNING id
	`
	var mealID int64
	err := db.DB.QueryRow(query,
		u.UserID,
		u.StrMeal,
		u.StrCategory,
		u.StrArea,
		u.StrInstructions,
		u.StrMealThumb,
		u.StrIngredient1,
		u.StrIngredient2,
		u.StrIngredient3,
		u.StrIngredient4,
		u.StrIngredient5,
		u.StrIngredient6,
		u.StrIngredient7,
		u.StrIngredient8,
		u.StrIngredient9,
		u.StrIngredient10,
		u.StrIngredient11,
		u.StrIngredient12,
		u.StrIngredient13,
		u.StrIngredient14,
		u.StrIngredient15,
		u.StrIngredient16,
		u.StrIngredient17,
		u.StrIngredient18,
		u.StrIngredient19,
		u.StrIngredient20,
		u.StrMeasure1,
		u.StrMeasure2,
		u.StrMeasure3,
		u.StrMeasure4,
		u.StrMeasure5,
		u.StrMeasure6,
		u.StrMeasure7,
		u.StrMeasure8,
		u.StrMeasure9,
		u.StrMeasure10,
		u.StrMeasure11,
		u.StrMeasure12,
		u.StrMeasure13,
		u.StrMeasure14,
		u.StrMeasure15,
		u.StrMeasure16,
		u.StrMeasure17,
		u.StrMeasure18,
		u.StrMeasure19,
		u.StrMeasure20).Scan(&mealID)

	if err != nil {
		return err
	}

	u.IDMeal = mealID
	return nil
}

//Returns all recipes for the specified user id
func GetAllUserRecipesByUserID(userId int64) ([]UserRecipe, error) {
	query := "SELECT * FROM user_recipes WHERE user_id = $1"
	rows, err := db.DB.Query(query, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var userRecipes []UserRecipe

	for rows.Next() {
		var userRecipe UserRecipe
		err := rows.Scan(
			&userRecipe.IDMeal,
			&userRecipe.UserID,
			&userRecipe.StrMeal,
			&userRecipe.StrCategory,
			&userRecipe.StrArea,
			&userRecipe.StrInstructions,
			&userRecipe.StrMealThumb,
			&userRecipe.StrIngredient1,
			&userRecipe.StrIngredient2,
			&userRecipe.StrIngredient3,
			&userRecipe.StrIngredient4,
			&userRecipe.StrIngredient5,
			&userRecipe.StrIngredient6,
			&userRecipe.StrIngredient7,
			&userRecipe.StrIngredient8,
			&userRecipe.StrIngredient9,
			&userRecipe.StrIngredient10,
			&userRecipe.StrIngredient11,
			&userRecipe.StrIngredient12,
			&userRecipe.StrIngredient13,
			&userRecipe.StrIngredient14,
			&userRecipe.StrIngredient15,
			&userRecipe.StrIngredient16,
			&userRecipe.StrIngredient17,
			&userRecipe.StrIngredient18,
			&userRecipe.StrIngredient19,
			&userRecipe.StrIngredient20,
			&userRecipe.StrMeasure1,
			&userRecipe.StrMeasure2,
			&userRecipe.StrMeasure3,
			&userRecipe.StrMeasure4,
			&userRecipe.StrMeasure5,
			&userRecipe.StrMeasure6,
			&userRecipe.StrMeasure7,
			&userRecipe.StrMeasure8,
			&userRecipe.StrMeasure9,
			&userRecipe.StrMeasure10,
			&userRecipe.StrMeasure11,
			&userRecipe.StrMeasure12,
			&userRecipe.StrMeasure13,
			&userRecipe.StrMeasure14,
			&userRecipe.StrMeasure15,
			&userRecipe.StrMeasure16,
			&userRecipe.StrMeasure17,
			&userRecipe.StrMeasure18,
			&userRecipe.StrMeasure19,
			&userRecipe.StrMeasure20)

		if err != nil {
			return nil, err
		}

		userRecipes = append(userRecipes, userRecipe)
	}

	return userRecipes, nil
}

//Returns a single user's recipe based on the recipe's id
func GetSingleUserRecipe(recipeId int64) (*UserRecipe, error) {
	query := "SELECT * FROM user_recipes WHERE id = $1"
	row := db.DB.QueryRow(query, recipeId)

	var userRecipe UserRecipe
	err := row.Scan(
		&userRecipe.IDMeal,
		&userRecipe.UserID,
		&userRecipe.StrMeal,
		&userRecipe.StrCategory,
		&userRecipe.StrArea,
		&userRecipe.StrInstructions,
		&userRecipe.StrMealThumb,
		&userRecipe.StrIngredient1,
		&userRecipe.StrIngredient2,
		&userRecipe.StrIngredient3,
		&userRecipe.StrIngredient4,
		&userRecipe.StrIngredient5,
		&userRecipe.StrIngredient6,
		&userRecipe.StrIngredient7,
		&userRecipe.StrIngredient8,
		&userRecipe.StrIngredient9,
		&userRecipe.StrIngredient10,
		&userRecipe.StrIngredient11,
		&userRecipe.StrIngredient12,
		&userRecipe.StrIngredient13,
		&userRecipe.StrIngredient14,
		&userRecipe.StrIngredient15,
		&userRecipe.StrIngredient16,
		&userRecipe.StrIngredient17,
		&userRecipe.StrIngredient18,
		&userRecipe.StrIngredient19,
		&userRecipe.StrIngredient20,
		&userRecipe.StrMeasure1,
		&userRecipe.StrMeasure2,
		&userRecipe.StrMeasure3,
		&userRecipe.StrMeasure4,
		&userRecipe.StrMeasure5,
		&userRecipe.StrMeasure6,
		&userRecipe.StrMeasure7,
		&userRecipe.StrMeasure8,
		&userRecipe.StrMeasure9,
		&userRecipe.StrMeasure10,
		&userRecipe.StrMeasure11,
		&userRecipe.StrMeasure12,
		&userRecipe.StrMeasure13,
		&userRecipe.StrMeasure14,
		&userRecipe.StrMeasure15,
		&userRecipe.StrMeasure16,
		&userRecipe.StrMeasure17,
		&userRecipe.StrMeasure18,
		&userRecipe.StrMeasure19,
		&userRecipe.StrMeasure20)

	if err != nil {
		return nil, err
	}

	return &userRecipe, nil
}

//Modification of a user's recipe
func (u UserRecipe) Update() error {
	query := `
	UPDATE user_recipes
	SET 
		user_id = $1,
		str_meal = $2,
		str_category = $3,
		str_area = $4,
		str_instructions = $5,
		str_meal_thumb = $6,
		str_ingredient1 = $7,
		str_ingredient2 = $8,
		str_ingredient3 = $9,
		str_ingredient4 = $10,
		str_ingredient5 = $11,
		str_ingredient6 = $12,
		str_ingredient7 = $13,
		str_ingredient8 = $14,
		str_ingredient9 = $15,
		str_ingredient10 = $16,
		str_ingredient11 = $17,
		str_ingredient12 = $18,
		str_ingredient13 = $19,
		str_ingredient14 = $20,
		str_ingredient15 = $21,
		str_ingredient16 = $22,
		str_ingredient17 = $23,
		str_ingredient18 = $24,
		str_ingredient19 = $25,
		str_ingredient20 = $26,
		str_measure1 = $27,
		str_measure2 = $28,
		str_measure3 = $29,
		str_measure4 = $30,
		str_measure5 = $31,
		str_measure6 = $32,
		str_measure7 = $33,
		str_measure8 = $34,
		str_measure9 = $35,
		str_measure10 = $36,
		str_measure11 = $37,
		str_measure12 = $38,
		str_measure13 = $39,
		str_measure14 = $40,
		str_measure15 = $41,
		str_measure16 = $42,
		str_measure17 = $43,
		str_measure18 = $44,
		str_measure19 = $45,
		str_measure20 = $46
	WHERE 
		id = $47
	`
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(
		u.UserID,
		u.StrMeal,
		u.StrCategory,
		u.StrArea,
		u.StrInstructions,
		u.StrMealThumb,
		u.StrIngredient1,
		u.StrIngredient2,
		u.StrIngredient3,
		u.StrIngredient4,
		u.StrIngredient5,
		u.StrIngredient6,
		u.StrIngredient7,
		u.StrIngredient8,
		u.StrIngredient9,
		u.StrIngredient10,
		u.StrIngredient11,
		u.StrIngredient12,
		u.StrIngredient13,
		u.StrIngredient14,
		u.StrIngredient15,
		u.StrIngredient16,
		u.StrIngredient17,
		u.StrIngredient18,
		u.StrIngredient19,
		u.StrIngredient20,
		u.StrMeasure1,
		u.StrMeasure2,
		u.StrMeasure3,
		u.StrMeasure4,
		u.StrMeasure5,
		u.StrMeasure6,
		u.StrMeasure7,
		u.StrMeasure8,
		u.StrMeasure9,
		u.StrMeasure10,
		u.StrMeasure11,
		u.StrMeasure12,
		u.StrMeasure13,
		u.StrMeasure14,
		u.StrMeasure15,
		u.StrMeasure16,
		u.StrMeasure17,
		u.StrMeasure18,
		u.StrMeasure19,
		u.StrMeasure20,
		u.IDMeal)
	return err
}

//Suppression of a user's recipe
func (u UserRecipe) Delete() error {
	query := "DELETE FROM user_recipes WHERE id = $1"
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(u.IDMeal)
	return err
}
