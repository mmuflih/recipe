package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

/**
 * Created by Muhammad Muflih Kholidin
 * https://github.com/mmuflih
 * muflic.24@gmail.com
**/

type P404Handler interface {
	Handle(c *gin.Context)
}

type p404H struct{}

func NewP404Handler() P404Handler {
	return &p404H{}
}

func (bh p404H) Handle(c *gin.Context) {
	data := "Sssssst! Silence is golden..."
	c.JSONP(http.StatusOK, data)
}
