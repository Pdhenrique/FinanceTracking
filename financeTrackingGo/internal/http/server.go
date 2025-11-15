package http

import (
	"context"
	"log"
	"net/http"
	"time"
)

type Server struct {
	server *http.Server
}

func NewServer(handler http.Handler, port string) *Server {
	return &Server{
		server: &http.Server{
			Addr:         ":" + port,
			Handler:      handler,
			ReadTimeout:  5 * time.Second,
			WriteTimeout: 55 * time.Second,
		},
	}
}

func (server *Server) Start() {
	go func() {
		log.Printf("Starting server on port %s", server.server.Addr)
		if err := server.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Printf("Failed to start server: %q ", err)
		}
	}()
}

func (server *Server) Stop() {
	log.Println("Stopping server...")
	cont, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	if err := server.server.Shutdown(cont); err != nil && err != http.ErrServerClosed {
		log.Printf("Could not shutdown in 60s: %v", err)
		return
	}
	log.Println("Server stopped")
}
