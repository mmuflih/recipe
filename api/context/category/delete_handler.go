package category

import (
	"errors"
	"time"
)

/**
 * Created by Muhammad Muflih Kholidin
 * https://github.com/mmuflih
 * muflic.24@gmail.com
 **/

func (mh mainH) Delete(id uint16) (interface{}, error) {
	cat, err := mh.repo.Find(id)
	if err != nil {
		return nil, errors.New("category data not found")
	}
	now := time.Now()
	cat.DeletedAt = &now
	err = mh.repo.Update(cat)
	if err != nil {
		return nil, errors.New("error while deleting category")
	}
	return "category deleted", nil
}
