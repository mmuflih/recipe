package recipe

import (
	"encoding/json"

	paginator "github.com/mmuflih/gorm-paginator"
	"github.com/mmuflih/recipe/api/http/core/request"
)

/**
 * Created by Muhammad Muflih Kholidin
 * https://github.com/mmuflih
 * muflic.24@gmail.com
 **/

func (h mainUC) List(req request.PaginationRequest) (interface{}, error) {
	paginate := h.recipeRepo.List(req.Query, req.Page, req.Size)

	return paginate, nil
}

func (h mainUC) ExtractInterface(source interface{}, result interface{}) error {
	dByte, err := json.Marshal(source)
	if err != nil {
		return err
	}
	err = json.Unmarshal(dByte, &result)
	if err != nil {
		return err
	}
	return nil
}

func (h mainUC) ExtractPaginator(paginate interface{}, result interface{}) error {
	lPaginator := paginator.Paginator{}
	err := h.ExtractInterface(paginate, &lPaginator)
	if err != nil {
		return err
	}
	err = h.ExtractInterface(lPaginator.Data, &result)
	if err != nil {
		return err
	}
	return nil
}
