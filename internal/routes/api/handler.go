package api

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type (
	Handler interface {
		GetOpenTicketCount(w http.ResponseWriter, r *http.Request)
	}

	handler struct {
		log     *slog.Logger
		service Service
	}
)

func NewHandler(log *slog.Logger, svc Service) Handler {
	return &handler{
		log:     log.With("Handler", "api"),
		service: svc,
	}
}

func Mount(r chi.Router, h Handler) {
	r.Route("/api", func(r chi.Router) {
		r.Route("/stats", func(r chi.Router) {
			r.Get("/open", h.GetOpenTicketCount)
		})
	})
}

func (h *handler) GetOpenTicketCount(w http.ResponseWriter, r *http.Request) {
	// TODO: Query database
	count, err := h.service.OpenTicketsCount(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	countString := fmt.Sprintf("%d", count)

	fmt.Fprintf(w, "%s", countString)
}
