package category

import "errors"

/**
 * Created by Muhammad Muflih Kholidin
 * https://github.com/mmuflih
 * muflic.24@gmail.com
 **/

func (h mainH) GetBySlug(slug string) (interface{}, error) {
	cat, err := h.repo.FindBySlug(slug)
	if err != nil {
		return nil, errors.New("category data not found")
	}
	return cat, nil
}
