package ingredient

import (
	"github.com/mmuflih/recipe/api/app/helper"
	"github.com/mmuflih/recipe/api/domain/repositories"
	"github.com/mmuflih/recipe/api/http/core/request"
	"github.com/mmuflih/recipe/api/http/requests"
)

/**
 * Created by Muhammad Muflih Kholidin
 * https://github.com/mmuflih
 * muflic.24@gmail.com
 **/

type MainUsecase interface {
	Add(req requests.IngredientRequest) (interface{}, error)
	Edit(id uint, req requests.IngredientRequest) (interface{}, error)
	GetBySlug(slug string) (interface{}, error)
	List(req request.PaginationRequest) (interface{}, error)
	Delete(id uint) (interface{}, error)
}

type mainUC struct {
	slugHelp helper.Slug
	repo     repositories.IngredientRepository
}

func NewMainUsecase(slugHelp helper.Slug,
	repo repositories.IngredientRepository) MainUsecase {
	return &mainUC{slugHelp, repo}
}
