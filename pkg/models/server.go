package models

import (
	"net/http"
	"time"
)

type Server struct {
	httpServer *http.Server
}

// Run the server with specific config properties
func (s *Server) Run(port string, handler http.Handler) error {
	// Server config properties
	s.httpServer = &http.Server{
		Addr:           ":" + port,
		Handler:        handler,
		MaxHeaderBytes: 1 << 20, // 1 MB
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}

	// Run the server
	return s.httpServer.ListenAndServe()
}
