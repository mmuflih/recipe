package container

import (
	"github.com/gin-gonic/gin"
	"github.com/mmuflih/recipe/api/http/handlers"
)

/**
 * Created by Muhammad Muflih Kholidin
 * https://github.com/mmuflih
 * muflic.24@gmail.com
**/

func InvokeRoute(route *gin.Engine,
	pingH handlers.PingHandler, p404H handlers.P404Handler,
	catH handlers.CategoryHandler, ingrH handlers.IngredientHandler,
	recipeH handlers.RecipeHandler,
) {
	/** not available route */
	route.NoRoute(p404H.Handle)

	/** ping */
	pingRoute := route.Group("/ping")
	pingRoute.GET("", pingH.Handle)

	/** v1 */
	v1 := route.Group("/v1")

	/** categories */
	catR := v1.Group("/categories")
	catR.POST("", catH.Add)
	catR.GET("", catH.List)
	catR.PUT("/:id", catH.Edit)
	catR.DELETE("/:id", catH.Delete)
	catR.GET("/slug/:slug", catH.GetBySlug)

	/** ingredients */
	ingR := v1.Group("/ingredients")
	ingR.POST("", ingrH.Add)
	ingR.GET("", ingrH.List)
	ingR.PUT("/:id", ingrH.Edit)
	ingR.DELETE("/:id", ingrH.Delete)
	ingR.GET("/slug/:slug", ingrH.GetBySlug)

	/** recipes */
	recipeR := v1.Group("/recipes")
	recipeR.POST("", recipeH.Add)
	recipeR.GET("", recipeH.List)
	recipeR.PUT("/:id", recipeH.Edit)
	recipeR.DELETE("/:id", recipeH.Delete)
	recipeR.GET("/slug/:slug", recipeH.GetBySlug)
}
