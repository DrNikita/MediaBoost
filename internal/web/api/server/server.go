package server

import (
	"net/http"

	"github.com/fatih/color"
	"github.com/gin-gonic/gin"
)

type Server struct {
	router *gin.Engine
}

func NewServer() *Server {
	gin.DefaultWriter = color.Output
	gin.DefaultErrorWriter = color.Error
	s := &Server{
		router: gin.Default(),
	}
	s.registerRoutes()
	return s
}

func (s *Server) registerRoutes() {
}

func (s *Server) Routes() gin.RoutesInfo {
	return s.router.Routes()
}

func (s *Server) ListenAndServe(addr string) error {
	return s.router.Run(addr)
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}
