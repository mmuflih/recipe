package category

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
	Add(req requests.CategoryRequest) (interface{}, error)
	Edit(id uint16, req requests.CategoryRequest) (interface{}, error)
	GetBySlug(slug string) (interface{}, error)
	List(req request.PaginationRequest) (interface{}, error)
	Delete(id uint16) (interface{}, error)
}

type mainH struct {
	slugHelp helper.Slug
	repo     repositories.CategoryRepository
}

func NewMainUsecase(slugHelp helper.Slug,
	repo repositories.CategoryRepository) MainUsecase {
	return &mainH{slugHelp, repo}
}
