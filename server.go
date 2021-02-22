package contact

import (
	"context"
	"net/http"
	"time"
)

// Server http server
type Server struct {
	server *http.Server
}

// Run server ...
func (s *Server) Run(port string, handler http.Handler) error{
	s.server = &http.Server{
		Addr: ":" + port,
		MaxHeaderBytes: 1 << 20,
		Handler: handler,
		ReadTimeout: 5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}
	
	return s.server.ListenAndServe()
}

// ShutDown server
func (s *Server) ShutDown(ctx context.Context) error {
	return s.server.Shutdown(ctx)
}