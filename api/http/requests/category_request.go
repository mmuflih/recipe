package requests

/**
 * Created by Muhammad Muflih Kholidin
 * https://github.com/mmuflih
 * muflic.24@gmail.com
**/

type CategoryRequest struct {
	Name string `json:"name" valid:"required"`
}
