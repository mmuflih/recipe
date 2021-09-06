package handlers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/mmuflih/recipe/api/context/category"
	"github.com/mmuflih/recipe/api/http/core/request"
	"github.com/mmuflih/recipe/api/http/core/response"
	"github.com/mmuflih/recipe/api/http/requests"
)

/**
 * Created by Muhammad Muflih Kholidin
 * https://github.com/mmuflih
 * muflic.24@gmail.com
**/

type CategoryHandler interface {
	Add(c *gin.Context)
	Edit(c *gin.Context)
	Delete(c *gin.Context)
	GetBySlug(c *gin.Context)
	List(c *gin.Context)
}

type categoryH struct {
	mainUsecase category.MainUsecase
	rr          request.Reader
}

func NewCategoryHandler(mainUsecase category.MainUsecase,
	rr request.Reader) CategoryHandler {
	return &categoryH{mainUsecase, rr}
}

func (ph *categoryH) Add(c *gin.Context) {
	req := requests.CategoryRequest{}
	err := request.RequestValidator(c, ph.rr, &req)
	if err != nil {
		response.Exception(c, err, 422)
		return
	}
	resp, err := ph.mainUsecase.Add(req)
	response.Json(c, resp, err)
}

func (ph *categoryH) Edit(c *gin.Context) {
	req := requests.CategoryRequest{}
	err := request.RequestValidator(c, ph.rr, &req)
	if err != nil {
		response.Exception(c, err, 422)
		return
	}
	id := ph.rr.GetRouteParamInt(c, "id")
	resp, err := ph.mainUsecase.Edit(uint16(id), req)
	response.Json(c, resp, err)
}

func (ph *categoryH) Delete(c *gin.Context) {
	id := ph.rr.GetRouteParamInt(c, "id")
	resp, err := ph.mainUsecase.Delete(uint16(id))
	response.Json(c, resp, err)
}

func (ph *categoryH) GetBySlug(c *gin.Context) {
	slug := ph.rr.GetRouteParam(c, "slug")
	fmt.Println("SLUG", slug)
	resp, err := ph.mainUsecase.GetBySlug(slug)
	response.Json(c, resp, err)
}

func (ph *categoryH) List(c *gin.Context) {
	req := request.PaginationFromRequest(ph.rr, c)
	resp, err := ph.mainUsecase.List(req)
	response.Json(c, resp, err)
}
