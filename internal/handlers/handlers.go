package handlers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"flexsupport/internal/models"
	"flexsupport/views/layouts"
	"flexsupport/views/pages"

	"github.com/go-chi/chi/v5"
)

// Handler holds dependencies for HTTP handlers
type Handler struct{}

// NewHandler creates a new Handler instance
func NewHandler() *Handler {
	return &Handler{}
}

// Dashboard renders the main dashboard view
func (h *Handler) Dashboard(w http.ResponseWriter, r *http.Request) {
	// TODO: Fetch real data from database
	fmt.Println("Rendering dashboard")
	page := pages.Dashboard(getMockTickets())
	err := layouts.BaseLayout(page).Render(r.Context(), w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// ListTickets handles the ticket listing page
func (h *Handler) ListTickets(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement filtering based on query parameters
	status := r.URL.Query().Get("status")
	search := r.URL.Query().Get("search")

	// TODO: Fetch real data from database
	tickets := getMockTickets()

	if status != "" {
		tickets = filterTicketsByStatus(tickets, status)
	}
	if search != "" {
		tickets = filterTicketsBySearch(tickets, search)
	}

	err := pages.TicketRows(tickets).Render(r.Context(), w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
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

// NewTicketForm renders the new ticket form
func (h *Handler) NewTicketForm(w http.ResponseWriter, r *http.Request) {
	page := pages.TicketForm(models.Ticket{})
	err := layouts.BaseLayout(page).Render(r.Context(), w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// CreateTicket handles ticket creation
func (h *Handler) CreateTicket(w http.ResponseWriter, r *http.Request) {
	// TODO: Parse form data
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Invalid form data", http.StatusBadRequest)
		return
	}

	// TODO: Validate and save to database
	log.Printf("Creating ticket: %+v", r.PostForm)

	// For now, redirect to dashboard
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

// ViewTicket renders the ticket detail view
func (h *Handler) ViewTicket(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ticket ID", http.StatusBadRequest)
		return
	}

	// TODO: Fetch from database
	ticket := getMockTicket(id)

	page := pages.TicketPage(ticket)
	err = layouts.BaseLayout(page).Render(r.Context(), w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	return
}

// EditTicketForm renders the edit ticket form
func (h *Handler) EditTicketForm(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ticket ID", http.StatusBadRequest)
		return
	}

	// TODO: Fetch from database
	ticket := getMockTicket(id)

	page := pages.TicketForm(ticket)
	err = layouts.BaseLayout(page).Render(r.Context(), w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	return
}

// UpdateTicket handles ticket updates
func (h *Handler) UpdateTicket(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Invalid form data", http.StatusBadRequest)
		return
	}

	// TODO: Update database
	log.Printf("Updating ticket %s: %+v", idStr, r.PostForm)

	http.Redirect(w, r, "/tickets/"+idStr, http.StatusSeeOther)
}

// SearchTickets handles ticket search (htmx endpoint)
func (h *Handler) SearchTickets(w http.ResponseWriter, r *http.Request) {
	search := r.URL.Query().Get("search")
	log.Printf("Searching tickets: %s", search)

	// TODO: Implement search
}

// UpdateTicketStatus handles status updates (htmx endpoint)
func (h *Handler) UpdateTicketStatus(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	status := r.FormValue("status")

	log.Printf("Updating ticket %s status to: %s", idStr, status)

	// TODO: Update database

	// Return updated status badge
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintf(w, `<span id="status-badge" class="px-3 py-1 inline-flex text-sm leading-5 font-semibold rounded-full bg-yellow-100 text-yellow-800">%s</span>`, status)
}

// AddPart adds a part to a ticket (htmx endpoint)
func (h *Handler) AddPart(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Invalid form data", http.StatusBadRequest)
		return
	}

	log.Printf("Adding part to ticket %s: %+v", idStr, r.PostForm)

	// TODO: Save to database

	// Return new part HTML
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintf(w, `<div class="flex items-center justify-between p-3 bg-gray-50 rounded-md">
		<div class="flex-1">
			<span class="text-sm font-medium text-gray-900">%s</span>
			<span class="text-sm text-gray-500 ml-2">× %s</span>
		</div>
		<div class="flex items-center gap-3">
			<span class="text-sm font-medium text-gray-900">$%s</span>
		</div>
	</div>`, r.FormValue("part_name"), r.FormValue("quantity"), r.FormValue("cost"))
}

// DeletePart removes a part from a ticket (htmx endpoint)
func (h *Handler) DeletePart(w http.ResponseWriter, r *http.Request) {
	ticketID := chi.URLParam(r, "id")
	partID := chi.URLParam(r, "partId")

	log.Printf("Deleting part %s from ticket %s", partID, ticketID)

	// TODO: Delete from database

	w.WriteHeader(http.StatusOK)
}

// AddNote adds a work note to a ticket (htmx endpoint)
func (h *Handler) AddNote(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	note := r.FormValue("note")

	log.Printf("Adding note to ticket %s: %s", idStr, note)

	// TODO: Save to database

	// Return new note HTML
	w.Header().Set("Content-Type", "text/html")
	now := time.Now().Format("Jan 2, 2006 3:04 PM")
	fmt.Fprintf(w, `<div class="p-3 bg-gray-50 rounded-md">
		<div class="flex justify-between items-start mb-1">
			<span class="text-sm font-medium text-gray-900">Current User</span>
			<span class="text-xs text-gray-500">%s</span>
		</div>
		<p class="text-sm text-gray-700 whitespace-pre-line">%s</p>
	</div>`, now, note)
}

// TechnicianTicketView shows the detailed technician view of a ticket

// GetOpenTicketsCount returns the count of open tickets (htmx endpoint)
func (h *Handler) GetOpenTicketsCount(w http.ResponseWriter, r *http.Request) {
	// TODO: Query database
	fmt.Fprintf(w, "12")
}

func getMockTickets() []models.Ticket {
	return []models.Ticket{
		{
			ID:               1001,
			Status:           "new",
			Priority:         "high",
			CustomerName:     "John Doe",
			CustomerPhone:    "(555) 123-4567",
			ItemType:         "Smartphone",
			ItemModel:        "iPhone 13 Pro",
			IssueDescription: "Cracked screen, needs replacement",
			AssignedTo:       "Mike Tech",
			DueDate:          timePtr(time.Now().Add(48 * time.Hour)),
		},
		{
			ID:               1002,
			Status:           "in_progress",
			Priority:         "normal",
			CustomerName:     "Jane Smith",
			CustomerPhone:    "(555) 987-6543",
			ItemType:         "Laptop",
			ItemModel:        "MacBook Pro 2020",
			IssueDescription: "Battery not charging",
			AssignedTo:       "Sarah Tech",
		},
	}
}

func getMockTicket(id int) models.Ticket {
	dueDate := time.Now().Add(48 * time.Hour)
	return models.Ticket{
		ID:               int64(id),
		Status:           "in_progress",
		Priority:         "high",
		CustomerName:     "John Doe",
		CustomerPhone:    "(555) 123-4567",
		CustomerEmail:    "john@example.com",
		ItemType:         "Smartphone",
		ItemBrand:        "Apple",
		ItemModel:        "iPhone 13 Pro",
		SerialNumber:     "ABC123456789",
		IssueDescription: "Screen is completely shattered after being dropped. Touch functionality still works but glass is unsafe.",
		EstimatedCost:    150.00,
		AssignedTo:       "Mike Tech",
		DueDate:          &dueDate,
		CreatedAt:        time.Now().Add(-24 * time.Hour),
		UpdatedAt:        time.Now(),
		CreatedBy:        "Front Desk",
		Parts: []models.Part{
			{ID: 1, Name: "iPhone 13 Pro Screen Assembly", Quantity: 1, Cost: 89.99},
			{ID: 2, Name: "Screen Adhesive", Quantity: 1, Cost: 5.99},
		},
		Notes: []models.WorkNote{
			{
				ID:        1,
				Author:    "Mike Tech",
				Content:   "Customer confirmed backup was done. Safe to proceed.",
				Timestamp: time.Now().Add(-2 * time.Hour),
			},
		},
		TotalPartsCost: 95.98,
	}
}

func getMockTechnicians() []models.Technician {
	return []models.Technician{
		{ID: 1, Name: "Mike Tech", ActiveJobs: 3, IsAvailable: true},
		{ID: 2, Name: "Sarah Tech", ActiveJobs: 2, IsAvailable: true},
		{ID: 3, Name: "Bob Repair", ActiveJobs: 5, IsAvailable: false},
	}
}

func timePtr(t time.Time) *time.Time {
	return &t
}
