package models

import "time"

/**
 * Created by Muhammad Muflih Kholidin
 * https://github.com/mmuflih
 * muflic.24@gmail.com
 **/

type RecipeDetail struct {
	ID           uint       `json:"id"`
	RecipeID     uint       `json:"recipe_id"`
	IngredientID uint       `json:"ingredient_id"`
	Qty          string     `json:"qty"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
	DeletedAt    *time.Time `json:"deleted_at"`
}

func NewRecipeDetail(recipeID, ingredientID uint, qty string) *RecipeDetail {
	return &RecipeDetail{
		RecipeID:     recipeID,
		IngredientID: ingredientID,
		Qty:          qty,
	}
}
