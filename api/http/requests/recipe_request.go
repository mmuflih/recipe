package requests

/**
 * Created by Muhammad Muflih Kholidin
 * https://github.com/mmuflih
 * muflic.24@gmail.com
**/

type RecipeRequest struct {
	Name        string         `json:"name" valid:"required"`
	CategoryID  uint16         `json:"category_id" valid:"required"`
	Ingredients []RecipeDetail `json:"ingredients" valid:"required"`
}

type RecipeDetail struct {
	IngredientID uint   `json:"ingredient_id" valid:"required"`
	Qty          string `json:"qty" valid:"required"`
}
