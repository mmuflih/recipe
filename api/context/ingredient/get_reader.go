package ingredient

import "errors"

/**
 * Created by Muhammad Muflih Kholidin
 * https://github.com/mmuflih
 * muflic.24@gmail.com
 **/

func (h mainUC) GetBySlug(slug string) (interface{}, error) {
	cat, err := h.repo.FindBySlug(slug)
	if err != nil {
		return nil, errors.New("ingredient data not found")
	}
	return cat, nil
}
