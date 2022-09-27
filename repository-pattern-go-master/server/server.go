package server

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/hyperyuri/repository-pattern-go/server/routes"
)

type Server struct {
	server *gin.Engine
}

func NewServer() Server {
	return Server{

		server: gin.Default(),
	}
}

func (s *Server) Run() {
	router := routes.ConfigRoutes(s.server)

	log.Fatal(router.Run(":8080"))

}
