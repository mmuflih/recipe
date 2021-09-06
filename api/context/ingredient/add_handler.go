package ingredient

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

func (h mainUC) Add(req requests.IngredientRequest) (interface{}, error) {
	ing := models.NewIngredient(req.Name, h.slugHelp.Make("ingredients", req.Name))

	err := h.repo.Save(ing)
	if err != nil {
		return nil, errors.New("error while saving ingredient")
	}
	return ing, nil
}
