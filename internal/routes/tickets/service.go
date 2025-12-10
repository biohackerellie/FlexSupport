package tickets

import (
	"context"
	"errors"
	"log/slog"
	"strings"
	"time"

	"flexsupport/internal/models"
)

type (
	Service interface {
		Search(ctx context.Context, search, status string) ([]models.Ticket, error)
		Get(ctx context.Context, id int64) (models.Ticket, error)
	}

	service struct {
		log *slog.Logger
	}
)

func NewService(log *slog.Logger) Service {
	return &service{
		log: log.With("Service", "Tickets"),
	}
}

func (s service) Search(ctx context.Context, search, status string) ([]models.Ticket, error) {
	s.log.Debug("Searching for tickets", "search", search, "status", status)
	tickets := getMockTickets()
	if status != "" {
		tickets = filterTicketsByStatus(tickets, status)
	}
	if search != "" {
		tickets = filterTicketsBySearch(tickets, search)
	}
	return tickets, nil
}

func (s service) Get(ctx context.Context, id int64) (models.Ticket, error) {
	tickets := getMockTickets()
	for _, ticket := range tickets {
		if ticket.ID == id {
			return ticket, nil
		}
	}
	return models.Ticket{}, errors.New("ticket not found")
}

/**
* Temporary mock data and filtering
* TODO: Fetch real data from database
 */

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

func getMockTicket(id int64) models.Ticket {
	return models.Ticket{
		ID:               id,
		Status:           "in_progress",
		Priority:         "high",
		CustomerName:     "John Doe",
		CustomerPhone:    "(555) 123-4567",
		ItemType:         models.Bag,
		ItemModel:        "Backpack",
		IssueDescription: "Broken strap, needs replacement",
		AssignedTo:       "Mike Tech",
		DueDate:          time.Now().Add(48 * time.Hour),
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
