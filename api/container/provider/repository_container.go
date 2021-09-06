package provider

import "github.com/mmuflih/recipe/api/domain/repositories"

/**
 * Created by Muhammad Muflih Kholidin
 * https://github.com/mmuflih
 * muflic.24@gmail.com
**/

func Repositories() []interface{} {
	var r []interface{}

	r = append(r, repositories.NewCategoryRepo)
	r = append(r, repositories.NewIngredientRepo)
	r = append(r, repositories.NewRecipeRepo)
	r = append(r, repositories.NewRecipeDetailRepo)

	return r
}
