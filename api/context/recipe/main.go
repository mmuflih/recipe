package recipe

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
	Add(req requests.RecipeRequest) (interface{}, error)
	Edit(id uint, req requests.RecipeRequest) (interface{}, error)
	GetBySlug(slug string) (interface{}, error)
	List(req request.PaginationRequest) (interface{}, error)
	Delete(id uint) (interface{}, error)
}

type mainUC struct {
	slugHelp      helper.Slug
	recipeRepo    repositories.RecipeRepository
	recipeDtlRepo repositories.RecipeDetailRepository
	ingRepo       repositories.IngredientRepository
	catRepo       repositories.CategoryRepository
}

func NewMainUsecase(slugHelp helper.Slug,
	recipeRepo repositories.RecipeRepository,
	recipeDtlRepo repositories.RecipeDetailRepository,
	ingRepo repositories.IngredientRepository,
	catRepo repositories.CategoryRepository) MainUsecase {
	return &mainUC{slugHelp, recipeRepo, recipeDtlRepo,
		ingRepo, catRepo}
}
