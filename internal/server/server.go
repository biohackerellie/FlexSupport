package server

import (
	"log/slog"
	"net/http"
)

type Server struct {
	router *http.ServeMux
	logger *slog.Logger
}

func NewServer(router *http.ServeMux, logger *slog.Logger) *Server {
	return &Server{
		router: router,
		logger: logger,
	}
}
