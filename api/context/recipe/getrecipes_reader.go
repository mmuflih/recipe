package recipe

import "github.com/mmuflih/recipe/api/http/core/request"

/**
 * Created by Muhammad Muflih Kholidin
 * https://github.com/mmuflih
 * muflic.24@gmail.com
 **/

func (h mainUC) List(req request.PaginationRequest) (interface{}, error) {
	return h.recipeRepo.List(req.Query, req.Page, req.Size), nil
}
