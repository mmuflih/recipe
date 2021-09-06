package request

import (
	"github.com/gin-gonic/gin"
)

/**
 * Created by Muhammad Muflih Kholidin
 * https://github.com/mmuflih
 * muflic.24@gmail.com
 **/

type PaginationRequest struct {
	Page  int    `json:"-"`
	Size  int    `json:"-"`
	Query string `json:"-"`
}

func PaginationFromRequest(rr Reader, c *gin.Context) PaginationRequest {
	r := c.Request
	page := rr.GetQueryInt(r, "page")
	if page == 0 {
		page = 1
	}
	size := rr.GetQueryInt(r, "size")
	if size == 0 {
		size = 100
	}
	pagReq := PaginationRequest{}
	rr.GetJsonData(r, &pagReq)
	pagReq.Page = page
	pagReq.Size = size
	pagReq.Query = rr.GetQuery(r, "q")

	return pagReq
}
