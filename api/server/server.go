package server

import (
	"github.com/labstack/echo/v4"
	"golang.org/x/net/http2"
	"time"
)

type MyServer struct {
	server *http2.Server
}

func NewServer() *MyServer {

	s := &http2.Server{
		MaxConcurrentStreams: 250,
		//MaxHandlers:          r,
		MaxReadFrameSize:     1048576,
		IdleTimeout:          10 * time.Second,
	}

	return &MyServer{s}
}

func (s MyServer) Run(r *echo.Echo) {
	r.Logger.Fatal(r.StartH2CServer( ":8080", s.server))
}
