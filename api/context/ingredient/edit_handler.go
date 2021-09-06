package ingredient

import (
	"errors"

	"github.com/mmuflih/recipe/api/http/requests"
)

/**
 * Created by Muhammad Muflih Kholidin
 * https://github.com/mmuflih
 * muflic.24@gmail.com
 **/

func (h mainUC) Edit(id uint, req requests.IngredientRequest) (interface{}, error) {
	cat, err := h.repo.Find(id)
	if err != nil {
		return nil, errors.New("ingredient data not found")
	}
	cat.Name = req.Name
	err = h.repo.Update(cat)
	if err != nil {
		return nil, errors.New("error while editing ingredient")
	}
	return cat, nil
}
