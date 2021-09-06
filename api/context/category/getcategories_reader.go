package category

import "github.com/mmuflih/recipe/api/http/core/request"

/**
 * Created by Muhammad Muflih Kholidin
 * https://github.com/mmuflih
 * muflic.24@gmail.com
 **/

func (h mainH) List(req request.PaginationRequest) (interface{}, error) {
	return h.repo.List(req.Query, req.Page, req.Size), nil
}
