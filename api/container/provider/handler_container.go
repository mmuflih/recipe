package provider

import "github.com/mmuflih/recipe/api/http/handlers"

/**
 * Created by Muhammad Muflih Kholidin
 * https://github.com/mmuflih
 * muflic.24@gmail.com
**/

func Handlers() []interface{} {
	var h []interface{}

	h = append(h, handlers.NewPingHandler)
	h = append(h, handlers.NewP404Handler)
	h = append(h, handlers.NewCategoryHandler)
	h = append(h, handlers.NewIngredientHandler)
	h = append(h, handlers.NewRecipeHandler)

	return h
}
