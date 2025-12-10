package dashboard

import (
	"log/slog"
	"net/http"
	"strings"

	"flexsupport/internal/layout"
	"flexsupport/internal/models"
	"flexsupport/internal/utils"

	"github.com/go-chi/chi/v5"
)

type (
	Handler interface {
		Get(w http.ResponseWriter, r *http.Request)
	}
	handler struct {
		log     *slog.Logger
		service Service
	}
)

func NewHandler(log *slog.Logger, svc Service) Handler {
	return &handler{
		log:     log,
		service: svc,
	}
}

func Mount(r chi.Router, h Handler) {
	r.Get("/", h.Get)
}

func (h handler) Get(w http.ResponseWriter, r *http.Request) {
	tickets, err := h.service.List(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	isMobile := utils.IsMobileUA(r.UserAgent())
	h.log.Debug("isMobile", "isMobile", isMobile)

	v := layout.Handler(Dashboard(tickets, isMobile))
	v.ServeHTTP(w, r)
}

func filterTicketsByStatus(tickets []models.Ticket, status string) []models.Ticket {
	result := make([]models.Ticket, 0)
	if status != "" {
		for _, ticket := range tickets {
			if ticket.Status.String() == status {
				result = append(result, ticket)
			}
		}
	}
	return result
}

func filterTicketsBySearch(tickets []models.Ticket, search string) []models.Ticket {
	result := make([]models.Ticket, 0)
	if search != "" {
		for _, ticket := range tickets {
			if strings.Contains(ticket.CustomerName, search) || strings.Contains(ticket.IssueDescription, search) {
				result = append(result, ticket)
			}
		}
	}
	return result
}
