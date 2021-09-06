package request

import (
	"errors"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
)

/**
 * Created by Muhammad Muflih Kholidin
 * https://github.com/mmuflih
 * muflic.24@gmail.com
 **/

func RequestValidator(c *gin.Context, rr Reader, req interface{}) error {
	err := rr.GetJsonData(c.Request, &req)
	if err != nil {
		return err
	}
	result, err := govalidator.ValidateStruct(req)
	if err != nil {
		return err
	}
	if !result {
		return errors.New("validate error")
	}
	return nil
}
