package model

import (
	"context"
	"net/http"
	"time"
)

type Server struct {
	httpServer *http.Server
}

func (s *Server) Run(port string, handler http.Handler) error {
	s.httpServer = &http.Server{
		Handler:        handler,
		Addr:           ":" + port,
		MaxHeaderBytes: 1 << 20, //1 Mb
		ReadTimeout:    time.Second * 10,
		WriteTimeout:   time.Second * 10,
	}
	return s.httpServer.ListenAndServe()
}

func (s *Server) ShutDown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
