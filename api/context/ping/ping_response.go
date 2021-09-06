package ping

/**
 * Created by Muhammad Muflih Kholidin
 * https://github.com/mmuflih
 * muflic.24@gmail.com
**/

type pingResponse struct {
	Pong string `json:"ping"`
}

func (this pingResponse) GetPong() string {
	return this.Pong
}
