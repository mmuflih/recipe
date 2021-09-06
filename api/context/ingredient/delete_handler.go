package ingredient

import (
	"errors"
	"time"
)

/**
 * Created by Muhammad Muflih Kholidin
 * https://github.com/mmuflih
 * muflic.24@gmail.com
 **/

func (mh mainUC) Delete(id uint) (interface{}, error) {
	cat, err := mh.repo.Find(id)
	if err != nil {
		return nil, errors.New("ingredient data not found")
	}
	now := time.Now()
	cat.DeletedAt = &now
	err = mh.repo.Update(cat)
	if err != nil {
		return nil, errors.New("error while deleting ingredient")
	}
	return "ingredient deleted", nil
}
