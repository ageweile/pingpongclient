package server

import (
	"github.com/gin-gonic/gin"
)

type Server struct {
}

func (s *Server) CreateRoutes() {

	engine := gin.Default()

	engine.POST("/calculation", handleCalculation)

	engine.Run(":8081")
}
