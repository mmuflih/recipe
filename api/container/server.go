package container

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mmuflih/envgo/conf"
)

/**
 * Created by Muhammad Muflih Kholidin
 * https://github.com/mmuflih
 * muflic.24@gmail.com
**/

type ServerRoute struct {
	config conf.Config
	router *gin.Engine
}

func NewRoute(c conf.Config, router *gin.Engine) *ServerRoute {
	return &ServerRoute{c, router}
}

func (s *ServerRoute) Run() {
	s.serverRun()
}

func (s *ServerRoute) serverRun() {
	log.Println("Application is running at ", time.Now().Format("2006-01-02 15:04:05.000"))
	port := s.config.GetString(`server.address`)
	s.router.Run(port)
}
