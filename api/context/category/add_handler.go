package category

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

func (h mainH) Add(req requests.CategoryRequest) (interface{}, error) {
	cat := models.NewCategory(req.Name, h.slugHelp.Make("categories", req.Name))

	err := h.repo.Save(cat)
	if err != nil {
		return nil, errors.New("error while saving category")
	}
	return cat, nil
}
