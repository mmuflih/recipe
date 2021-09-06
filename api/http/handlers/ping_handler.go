package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/mmuflih/recipe/api/context/ping"
	"github.com/mmuflih/recipe/api/http/core/response"
	"github.com/mmuflih/recipe/api/http/requests"
)

/**
 * Created by Muhammad Muflih Kholidin
 * https://github.com/mmuflih
 * muflic.24@gmail.com
**/

type PingHandler interface {
	Handle(c *gin.Context)
}

type pingH struct {
	puc ping.PingUsecase
}

func NewPingHandler(puc ping.PingUsecase) PingHandler {
	return &pingH{puc}
}

func (ph *pingH) Handle(c *gin.Context) {
	req := requests.PingRequest{}
	resp, err := ph.puc.Ping(req)
	response.Json(c, resp, err)
}
