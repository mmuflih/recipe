package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/mmuflih/recipe/api/context/recipe"
	"github.com/mmuflih/recipe/api/http/core/request"
	"github.com/mmuflih/recipe/api/http/core/response"
	"github.com/mmuflih/recipe/api/http/requests"
)

/**
 * Created by Muhammad Muflih Kholidin
 * https://github.com/mmuflih
 * muflic.24@gmail.com
**/

type RecipeHandler interface {
	Add(c *gin.Context)
	Edit(c *gin.Context)
	Delete(c *gin.Context)
	GetBySlug(c *gin.Context)
	List(c *gin.Context)
}

type recipeH struct {
	mainUsecase recipe.MainUsecase
	rr          request.Reader
}

func NewRecipeHandler(mainUsecase recipe.MainUsecase,
	rr request.Reader) RecipeHandler {
	return &recipeH{mainUsecase, rr}
}

func (ph *recipeH) Add(c *gin.Context) {
	req := requests.RecipeRequest{}
	err := request.RequestValidator(c, ph.rr, &req)
	if err != nil {
		response.Exception(c, err, 422)
		return
	}
	resp, err := ph.mainUsecase.Add(req)
	response.Json(c, resp, err)
}

func (ph *recipeH) Edit(c *gin.Context) {
	req := requests.RecipeRequest{}
	err := request.RequestValidator(c, ph.rr, &req)
	if err != nil {
		response.Exception(c, err, 422)
		return
	}
	id := ph.rr.GetRouteParamInt(c, "id")
	resp, err := ph.mainUsecase.Edit(uint(id), req)
	response.Json(c, resp, err)
}

func (ph *recipeH) Delete(c *gin.Context) {
	id := ph.rr.GetRouteParamInt(c, "id")
	resp, err := ph.mainUsecase.Delete(uint(id))
	response.Json(c, resp, err)
}

func (ph *recipeH) GetBySlug(c *gin.Context) {
	slug := ph.rr.GetRouteParam(c, "slug")
	resp, err := ph.mainUsecase.GetBySlug(slug)
	response.Json(c, resp, err)
}

func (ph *recipeH) List(c *gin.Context) {
	req := request.PaginationFromRequest(ph.rr, c)
	resp, err := ph.mainUsecase.List(req)
	response.Json(c, resp, err)
}
