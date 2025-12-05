package router

import (
	"flexsupport/internal/handlers"
	mw "flexsupport/internal/middleware"
	"flexsupport/static"
	"net/http"

	// "net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func NewRouter(h *handlers.Handler) *chi.Mux {

	r := chi.NewMux()
	// Dashboard

	r.Handle("/assets/*", http.StripPrefix("/assets/", static.AssetRouter()))
	// Tickets
	r.Group(func(r chi.Router) {
	r.Use(
		middleware.Logger,
		middleware.Recoverer,
		mw.CSPMiddleware,
		mw.TextHTMLMiddleware,
	)
	r.Get("/", h.Dashboard)
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
})

	// API endpoints (for htmx)
	r.Route("/api", func(r chi.Router) {
		r.Get("/stats/open", h.GetOpenTicketsCount)
	})


	return r
}
