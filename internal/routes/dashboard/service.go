package dashboard

import (
	"context"
	"log/slog"
	"time"

	"flexsupport/internal/models"
)

type (
	Service interface {
		List(ctx context.Context) ([]models.Ticket, error)
	}

	service struct {
		log *slog.Logger
	}
)

func NewService(log *slog.Logger) Service {
	return &service{
		log: log.With("Dashboard Service"),
	}
}

func (s service) List(context.Context) ([]models.Ticket, error) {
	return getMockTickets(), nil
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
