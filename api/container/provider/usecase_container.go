package provider

import (
	"github.com/mmuflih/recipe/api/context/category"
	"github.com/mmuflih/recipe/api/context/ingredient"
	"github.com/mmuflih/recipe/api/context/ping"
	"github.com/mmuflih/recipe/api/context/recipe"
)

/**
 * Created by Muhammad Muflih Kholidin
 * https://github.com/mmuflih
 * muflic.24@gmail.com
**/

func Usecases() []interface{} {

	var u []interface{}

	u = append(u, ping.NewPingUsecase)
	u = append(u, category.NewMainUsecase)
	u = append(u, ingredient.NewMainUsecase)
	u = append(u, recipe.NewMainUsecase)

	return u
}
