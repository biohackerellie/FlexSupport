package tickets

import (
	"log/slog"
	"net/http"

	"flexsupport/internal/layout"
	"flexsupport/internal/models"
)

type Handler struct {
	log *slog.Logger
}

func (h *Handler) NewTicketForm(w http.ResponseWriter, r *http.Request) {
	page := TicketForm(models.Ticket{})
	err := layout.BaseLayout(page).Render(r.Context(), w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
