package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/realHoangHai/go-assmin/config"
	_ "github.com/realHoangHai/go-assmin/docs/swagger"
	"github.com/realHoangHai/go-assmin/internal/middleware"
	"github.com/realHoangHai/go-assmin/internal/service"
	"github.com/realHoangHai/go-assmin/pkg/log"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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
	r := gin.New()
	r.Use(s.handler.Recover())

	r.GET("/", s.service.Welcome())
	r.POST("/api/register", s.service.CreateUser())
	r.POST("/api/login", s.service.Login())
	r.POST("/api/renew-token", s.service.RenewToken())

	swaggerURL := ginSwagger.URL(fmt.Sprintf("0.0.0.0%s/swagger/doc.json", addr)) // the  url poiting to API definition
	r.GET("swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, swaggerURL))
	log.Infof(fmt.Sprintf("Starting server on http://localhost%s", addr))

	// start server
	go func() {
		if err := r.Run(addr); err != nil {
			log.Fatalf("Could not start server: %v", err)
		}
	}()
}
