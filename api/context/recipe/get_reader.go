package recipe

import (
	"errors"
	"fmt"

	"github.com/mmuflih/recipe/api/domain/models"
)

/**
 * Created by Muhammad Muflih Kholidin
 * https://github.com/mmuflih
 * muflic.24@gmail.com
 **/

func (h mainUC) GetBySlug(slug string) (interface{}, error) {
	recipe, err := h.recipeRepo.FindBySlug(slug)
	if err != nil {
		return nil, errors.New("recipe data not found")
	}

	details, _ := h.recipeDtlRepo.FindByRecipe(recipe.ID)
	return h.createResponse(recipe, details)
}

func (h mainUC) createResponse(recipe *models.Recipe,
	details []*models.RecipeDetail) (interface{}, error) {

	ingredients := make([]ingredientResponse, 0)
	for _, dtl := range details {
		ing, err := h.ingRepo.Find(dtl.IngredientID)
		if err != nil {
			fmt.Println("error ingredient =>", err)
			continue
		}
		ingredients = append(ingredients, ingredientResponse{
			Ingredient: ing.PublicResponse(),
			Qty:        dtl.Qty,
		})
	}
	return detailResponse{
		Recipe:      recipe,
		Ingredients: ingredients,
	}, nil
}

type detailResponse struct {
	*models.Recipe
	Ingredients []ingredientResponse `json:"ingredients"`
}

type ingredientResponse struct {
	Ingredient interface{} `json:"item"`
	Qty        string      `json:"qty"`
}
