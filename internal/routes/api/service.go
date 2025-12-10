package api

import (
	"context"
	"log/slog"
)

type (
	Service interface {
		OpenTicketsCount(ctx context.Context) (int, error)
	}
	service struct {
		log *slog.Logger
	}
)

func NewService(log *slog.Logger) Service {
	return &service{
		log: log.With("Service", "api"),
	}
}

func (s service) OpenTicketsCount(ctx context.Context) (int, error) {
	return 2, nil
}
