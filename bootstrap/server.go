package bootstrap

import (
	"arquitectura/domain/model"
	"arquitectura/infrastructure/handler"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
)

type Server struct {
	httpAddr string
	Engine   *gin.Engine
}

func New(host string, port uint, db *gorm.DB) *Server {

	srv := &Server{
		httpAddr: fmt.Sprintf("%s:%d", host, port),
		Engine:   gin.New(),
	}

	srv.Engine.Use(cors.New(cors.Config{
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowHeaders: []string{"Accept", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization", "Content-Type", "Apikey",
			"Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Access-Control-Allow-Methods", "Origin", "x-ibm-client-id", "X-Authorization-Type", "X-System-Type"},
		AllowCredentials: true,
	}))

	srv.Engine.Use(func (c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Content-Type", "application/json")
		c.Header("Strict-Transport-Security", "max-age=86400; includeSubDomains")
		c.Header("X-Content-Type-Options", "nosniff")
		c.Header("X-XSS-Protection", "1; mode=block")
	})

	routerSpecification := model.RouterSpecification{
		Api: srv.Engine,
		DB:  db,
	}
	handler.InitRoutes(routerSpecification)
	return srv
}

func (s *Server) Run() error {
	log.Println("Server running on", s.httpAddr)
	err := s.Engine.Run(s.httpAddr)
	if err != nil {
		log.Printf("Error starting server: %v", err)
		return err
	}

	return nil
}
