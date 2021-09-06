package recipe

import (
	"errors"

	"github.com/mmuflih/recipe/api/domain/models"
	"github.com/mmuflih/recipe/api/http/requests"
)

/**
 * Created by Muhammad Muflih Kholidin
 * https://github.com/mmuflih
 * muflic.24@gmail.com
 **/

func (h mainUC) Add(req requests.RecipeRequest) (interface{}, error) {
	/** validate category */
	_, err := h.catRepo.Find(req.CategoryID)
	if err != nil {
		return nil, errors.New("selected category not found")
	}

	recipe := models.NewRecipe(req.Name, h.slugHelp.Make("recipes", req.Name),
		req.CategoryID)

	tx := h.recipeDtlRepo.DBConn().Begin()
	err = h.recipeRepo.Save(recipe, tx)
	if err != nil {
		tx.Rollback()
		return nil, errors.New("error while saving recipe")
	}

	var details []*models.RecipeDetail
	for _, item := range req.Ingredients {
		_, err := h.ingRepo.Find(item.IngredientID)
		if err != nil {
			tx.Rollback()
			return nil, errors.New("selected ingredient not found")
		}
		dtl := models.NewRecipeDetail(recipe.ID, item.IngredientID, item.Qty)
		err = h.recipeDtlRepo.Save(dtl, tx)
		if err != nil {
			return nil, errors.New("error while saving recipe item")
		}
		details = append(details, dtl)
	}

	tx.Commit()

	return h.createResponse(recipe, details)
}
