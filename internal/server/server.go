package server

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type Server struct {
	router *gin.Engine
}

func New() *Server {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	srv := Server{
		router: router,
	}

	return &srv
}

func (s *Server) Run() error {
	log.Println("Server is running on http://localhost:8080")
	return s.router.Run(":8080")
}

func (s *Server) routes() {
	s.router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "healthy"})
	})
}
