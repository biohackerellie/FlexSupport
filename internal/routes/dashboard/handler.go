package dashboard

import (
	"log/slog"
	"net/http"
	"strings"
	"time"

	"flexsupport/internal/layout"
	"flexsupport/internal/models"

	"github.com/a-h/templ"
)

type Handler struct {
	log *slog.Logger
}

func NewHandler(log *slog.Logger) *Handler {
	return &Handler{
		log: log,
	}
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.Get(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
	status := r.URL.Query().Get("status")
	search := r.URL.Query().Get("search")
	template := r.URL.Query().Get("template")

	// TODO: Fetch real data from database
	tickets := getMockTickets()

	if status != "" {
		tickets = filterTicketsByStatus(tickets, status)
	}
	if search != "" {
		tickets = filterTicketsBySearch(tickets, search)
	}

	var opts []func(*templ.ComponentHandler)
	if template != "" {
		opts = append(opts, templ.WithFragments(template))
	}
	v := layout.Handler(Dashboard(tickets), opts...)
	v.ServeHTTP(w, r)
}

func getMockTickets() []models.Ticket {
	return []models.Ticket{
		{
			ID:               1001,
			Status:           "new",
			Priority:         "high",
			CustomerName:     "John Doe",
			CustomerPhone:    "(555) 123-4567",
			ItemType:         models.Bag,
			ItemModel:        "Backpack",
			IssueDescription: "Broken strap, needs replacement",
			AssignedTo:       "Mike Tech",
			DueDate:          time.Now().Add(48 * time.Hour),
		},
		{
			ID:               1002,
			Status:           "in_progress",
			Priority:         "normal",
			CustomerName:     "Jane Smith",
			CustomerPhone:    "(555) 987-6543",
			ItemType:         models.Boot,
			ItemModel:        "Keen",
			IssueDescription: "Sole needs replaced",
			AssignedTo:       "Sarah Tech",
			DueDate:          time.Now().Add(72 * time.Hour),
		},
	}
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
