package recipe

import (
	"errors"
	"time"

	"github.com/mmuflih/recipe/api/domain/models"
	"github.com/mmuflih/recipe/api/http/requests"
	"gorm.io/gorm"
)

/**
 * Created by Muhammad Muflih Kholidin
 * https://github.com/mmuflih
 * muflic.24@gmail.com
 **/

func (mu mainUC) Edit(id uint, req requests.RecipeRequest) (interface{}, error) {
	recipe, err := mu.recipeRepo.Find(id)
	if err != nil {
		return nil, errors.New("recipe data not found")
	}
	recipe.Name = req.Name
	recipe.CategoryID = req.CategoryID

	tx := mu.recipeRepo.DBConn().Begin()
	err = mu.recipeRepo.Update(recipe, tx)
	if err != nil {
		return nil, errors.New("error while editing recipe")
	}

	activeItems, err := mu.saveDetail(recipe.ID, req.Ingredients, tx)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	details, _ := mu.recipeDtlRepo.FindByRecipe(recipe.ID)
	err = mu.removeInactiveDetail(details, activeItems, tx)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()

	return mu.createResponse(recipe, details)
}

func (mu mainUC) saveDetail(recipeID uint, ingredients []requests.RecipeDetail,
	tx *gorm.DB) (map[uint]*models.RecipeDetail, error) {

	activeItems := make(map[uint]*models.RecipeDetail)
	for _, item := range ingredients {
		_, err := mu.ingRepo.Find(item.IngredientID)
		if err != nil {
			return activeItems, errors.New("selected ingredient not found")
		}

		dtl, err := mu.recipeDtlRepo.FindByIngredient(recipeID, item.IngredientID)
		if err != nil {
			dtl = models.NewRecipeDetail(recipeID, item.IngredientID, item.Qty)
			err = mu.recipeDtlRepo.Save(dtl, tx)
			if err != nil {
				return activeItems, errors.New("error while saving recipe detail")
			}
			continue
		}
		dtl.Qty = item.Qty
		err = mu.recipeDtlRepo.Update(dtl, tx)
		if err != nil {
			return activeItems, errors.New("error while saving recipe detail")
		}

		activeItems[dtl.ID] = dtl
	}
	return activeItems, nil
}

func (mu mainUC) removeInactiveDetail(details []*models.RecipeDetail,
	activeItems map[uint]*models.RecipeDetail, tx *gorm.DB) error {
	var err error
	for _, dtl := range details {
		if activeItems[dtl.ID] == nil {
			now := time.Now()
			dtl.DeletedAt = &now
			err = mu.recipeDtlRepo.Update(dtl, tx)
		}

		if err != nil {
			return errors.New("error while deleting recipe detail")
		}
	}
	return nil
}
