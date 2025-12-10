package tickets

import (
	"log/slog"
	"net/http"
	"strconv"

	"flexsupport/internal/layout"
	"flexsupport/internal/models"
	"flexsupport/ui/partials/rows"

	"github.com/go-chi/chi/v5"
)

type (
	Handler interface {
		Search(w http.ResponseWriter, r *http.Request)
		Get(w http.ResponseWriter, r *http.Request)
		New(w http.ResponseWriter, r *http.Request)
	}

	handler struct {
		log     *slog.Logger
		service Service
	}
)

func NewHandler(log *slog.Logger, svc Service) Handler {
	return &handler{
		log:     log.With("Tickets Handler"),
		service: svc,
	}
}

func Mount(r chi.Router, h Handler) {
	r.Route("/tickets", func(r chi.Router) {
		r.Get("/", h.Search)
		r.Route("/{ticketId}", func(r chi.Router) {
			r.Get("/", h.Get)
		})
		r.Get("/new", h.New)
	})
}

func (h handler) Search(w http.ResponseWriter, r *http.Request) {
	status := r.URL.Query().Get("status")
	search := r.URL.Query().Get("search")
	tickets, err := h.service.Search(r.Context(), search, status)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	switch isHTMX(r) {
	case true:
		err = rows.TicketRows(tickets).Render(r.Context(), w)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	default:
		err = layout.BaseLayout(TicketsPage(tickets, search, status)).Render(r.Context(), w)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func (h handler) Get(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "ticketId")
	ticketID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		http.Error(w, "Invalid ticket ID", http.StatusBadRequest)
		return
	}
	ticket, err := h.service.Get(r.Context(), ticketID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	page := TicketPage(ticket)
	err = layout.BaseLayout(page).Render(r.Context(), w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h handler) New(w http.ResponseWriter, r *http.Request) {
	page := TicketForm(models.Ticket{})
	err := layout.BaseLayout(page).Render(r.Context(), w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func isHTMX(r *http.Request) bool {
	// Check for "HX-Request" header
	if r.Header.Get("HX-Request") != "" {
		return true
	}

	return false
}
