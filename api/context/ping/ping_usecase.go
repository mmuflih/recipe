package ping

/**
 * Created by Muhammad Muflih Kholidin
 * https://github.com/mmuflih
 * muflic.24@gmail.com
**/

type PingUsecase interface {
	Ping(PingRequest) (PingResponse, error)
}

type PingRequest interface {
}

type PingResponse interface {
	GetPong() string
}

type pingUsecase struct{}

func NewPingUsecase() PingUsecase {
	return &pingUsecase{}
}

func (this *pingUsecase) Ping(req PingRequest) (PingResponse, error) {
	resp := pingResponse{"Pong"}
	return resp, nil
}
