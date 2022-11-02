package server

import "github.com/gin-gonic/gin"

type Server struct {
}

func (s *Server) Start() {
	gin.SetMode(gin.DebugMode)
	r := gin.Default()
	r.Use(gin.Recovery())
}
