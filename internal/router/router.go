package router

import (
	"flexsupport/internal/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func NewRouter(h *handlers.Handler) *chi.Mux {

	r := chi.NewMux()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.StripSlashes)
	// Dashboard
	r.Get("/", h.Dashboard)

	// Tickets
	r.Route("/tickets", func(r chi.Router) {
		r.Get("/", h.ListTickets)
		r.Get("/new", h.NewTicketForm)
		r.Post("/", h.CreateTicket)
		r.Get("/search", h.SearchTickets)
		r.Get("/{id}", h.ViewTicket)
		r.Get("/{id}/edit", h.EditTicketForm)
		r.Post("/{id}", h.UpdateTicket)

		// Ticket actions
		r.Post("/{id}/status", h.UpdateTicketStatus)
		r.Post("/{id}/parts", h.AddPart)
		r.Delete("/{id}/parts/{partId}", h.DeletePart)
		r.Post("/{id}/notes", h.AddNote)
	})

	// Technician view
	r.Get("/technician", h.TechnicianQueue)
	r.Get("/technician/{id}", h.TechnicianTicketView)

	// API endpoints (for htmx)
	r.Route("/api", func(r chi.Router) {
		r.Get("/stats/open", h.GetOpenTicketsCount)
	})

	return r
}
