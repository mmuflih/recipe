package recipe

import (
	"errors"
	"time"
)

/**
 * Created by Muhammad Muflih Kholidin
 * https://github.com/mmuflih
 * muflic.24@gmail.com
 **/

func (mu mainUC) Delete(id uint) (interface{}, error) {
	cat, err := mu.recipeRepo.Find(id)
	if err != nil {
		return nil, errors.New("recipe data not found")
	}
	now := time.Now()
	cat.DeletedAt = &now

	tx := mu.recipeRepo.DBConn().Begin()
	err = mu.recipeRepo.Update(cat, tx)
	if err != nil {
		tx.Rollback()
		return nil, errors.New("error while deleting recipe")
	}

	tx.Commit()

	return "recipe deleted", err
}
