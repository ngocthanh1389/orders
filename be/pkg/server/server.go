package server

import (
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
)

type Server struct {
	e        *gin.Engine
	bindAddr string
}

func NewServer(addr string) *Server {
	engine := gin.New()
	engine.Use(gin.Recovery())
	gin.SetMode(gin.ReleaseMode)

	return &Server{
		bindAddr: addr,
		e:        engine,
	}
}

func (s *Server) Register(method string, path string, handler gin.HandlerFunc) {
	s.e.Handle(method, path, handler)
}

func (s *Server) Run() error {
	s.registerDebug()
	return s.e.Run(s.bindAddr)
}

func (s *Server) registerDebug() {
	pprof.Register(s.e, "/debug")
}
