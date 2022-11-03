package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/realHoangHai/go-assmin/config"
	"github.com/realHoangHai/go-assmin/internal/middleware"
	"github.com/realHoangHai/go-assmin/internal/service"
	"github.com/realHoangHai/go-assmin/pkg/log"
)

type Server struct {
	handler *middleware.Handler
	service *service.Service
}

func NewServer(handler *middleware.Handler, service *service.Service) (s *Server, err error) {
	s = &Server{
		handler: handler,
		service: service,
	}
	return s, err
}

func (s *Server) Start() {
	addr := fmt.Sprintf(":%s", config.C.Core.Port)

	gin.SetMode(config.C.Core.Enviroment)
	r := gin.Default()
	r.Use(s.handler.Recover())

	log.Infof(fmt.Sprintf("Starting server on http://localhost%s", addr))

	// start server
	go func() {
		if err := r.Run(addr); err != nil {
			log.Fatalf("Could not start server: %v", err)
		}
	}()
}
