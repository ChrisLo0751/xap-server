package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HttpServer struct {
	engine *gin.Engine
}

type HttpServerHandler func()

func NewHttpServer() *HttpServer {
	r := gin.Default()

	return &HttpServer{engine: r}
}

func (s *HttpServer) Run() {
	s.engine.Run(":8080")
}

func (s *HttpServer) Register(path string, HttpServerHandler ...HttpServerHandler) error {
	return nil
}

func (s *HttpServer) SetRouters() {
	s.engine.GET("/api/v1/users/:username", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
}
