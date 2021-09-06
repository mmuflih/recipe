package category

import (
	"errors"

	"github.com/mmuflih/recipe/api/http/requests"
)

/**
 * Created by Muhammad Muflih Kholidin
 * https://github.com/mmuflih
 * muflic.24@gmail.com
 **/

func (h mainH) Edit(id uint16, req requests.CategoryRequest) (interface{}, error) {
	cat, err := h.repo.Find(id)
	if err != nil {
		return nil, errors.New("category data not found")
	}
	cat.Name = req.Name
	err = h.repo.Update(cat)
	if err != nil {
		return nil, errors.New("error while editing category")
	}
	return cat, nil
}
