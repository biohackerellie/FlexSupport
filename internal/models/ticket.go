package models

import "time"

type Status string

func (s Status) String() string {
	return string(s)
}

const (
	StatusNew          Status = "new"
	StatusInProgress   Status = "in_progress"
	StatusWaitingParts Status = "waiting_parts"
	StatusReady        Status = "ready"
	StatusCompleted    Status = "completed"
)

type Priority string

func (p Priority) String() string {
	return string(p)
}

const (
	PriorityLow    Priority = "low"
	PriorityNormal Priority = "normal"
	PriorityHigh   Priority = "high"
	PriorityUrgent Priority = "urgent"
)

type ItemType string

const (
	Boot  ItemType = "boot"
	Shoe  ItemType = "shoe"
	Bag   ItemType = "bag"
	Other ItemType = "other"
)

// Ticket represents a repair ticket in the system
type Ticket struct {
	ID          int64    `db:"id" json:"id"`
	Status      Status   `db:"status" json:"status"`
	Priority    Priority `db:"priority" json:"priority"` // low, normal, high, urgent
	ExternalTag string   `db:"external_tag" json:"external_tag"`

	// Customer information
	CustomerName  string `db:"customer_name" json:"customer_name"`
	CustomerPhone string `db:"customer_phone" json:"customer_phone"`
	CustomerEmail string `db:"customer_email" json:"customer_email"`

	// Item information
	ItemType     ItemType `db:"item_type" json:"item_type"`
	ItemBrand    string   `db:"item_brand" json:"item_brand"`
	ItemModel    string   `db:"item_model" json:"item_model"`
	SerialNumber string   `db:"serial_number" json:"serial_number"`

	// Repair details
	IssueDescription string  `db:"issue_description" json:"issue_description"`
	InternalNotes    string  `db:"internal_notes" json:"internal_notes"`
	EstimatedCost    float64 `db:"estimated_cost" json:"estimated_cost"`

	// Assignment and scheduling
	AssignedTo string    `db:"assigned_to" json:"assigned_to"`
	DueDate    time.Time `db:"due_date" json:"due_date"`

	// Metadata
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
	CreatedBy string    `db:"created_by" json:"created_by"`

	// Related data (loaded via joins)
	Parts          []Part     `db:"-" json:"parts,omitempty"`
	Notes          []WorkNote `db:"-" json:"notes,omitempty"`
	TotalPartsCost float64    `db:"-" json:"total_parts_cost"`
}

// Part represents a replacement part or material used in a repair
type Part struct {
	ID       int64     `json:"id"`
	TicketID int64     `json:"ticket_id"`
	Name     string    `json:"name"`
	Quantity int       `json:"quantity"`
	Cost     float64   `json:"cost"`
	AddedAt  time.Time `json:"added_at"`
	AddedBy  string    `json:"added_by"`
}

// WorkNote represents a work log entry or note on a ticket
type WorkNote struct {
	ID        int64     `json:"id"`
	TicketID  int64     `json:"ticket_id"`
	Content   string    `json:"content"`
	Author    string    `json:"author"`
	Timestamp time.Time `json:"timestamp"`
}

// Customer represents customer information (for future use)
type Customer struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name"`
	Phone       string    `json:"phone"`
	Email       string    `json:"email"`
	CreatedAt   time.Time `json:"created_at"`
	TotalOrders int       `json:"total_orders"`
}

// Technician represents a repair technician user
type Technician struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	ActiveJobs  int    `json:"active_jobs"`
	IsAvailable bool   `json:"is_available"`
}

// TicketStats represents dashboard statistics
type TicketStats struct {
	OpenTickets    int `json:"open_tickets"`
	InProgress     int `json:"in_progress"`
	Overdue        int `json:"overdue"`
	CompletedToday int `json:"completed_today"`
}

// StatusClass returns the Tailwind CSS class for the ticket status badge
func (t *Ticket) StatusClass() string {
	switch t.Status {
	case StatusNew:
		return "bg-blue-100 text-blue-800"
	case StatusInProgress:
		return "bg-yellow-100 text-yellow-800"
	case StatusWaitingParts:
		return "bg-orange-100 text-orange-800"
	case StatusReady:
		return "bg-green-100 text-green-800"
	case StatusCompleted:
		return "bg-gray-100 text-gray-800"
	default:
		return "bg-gray-100 text-gray-800"
	}
}

// StatusDisplay returns a human-readable status string
func (t *Ticket) StatusDisplay() string {
	switch t.Status {
	case StatusNew:
		return "New"
	case StatusInProgress:
		return "In Progress"
	case StatusWaitingParts:
		return "Waiting for Parts"
	case StatusReady:
		return "Ready for Pickup"
	case StatusCompleted:
		return "Completed"
	default:
		return t.Status.String()
	}
}

// TotalCost calculates the total cost including parts and estimated labor
func (t *Ticket) TotalCost() float64 {
	return t.EstimatedCost + t.TotalPartsCost
}

// IsOverdue checks if the ticket is past its due date
func (t *Ticket) IsOverdue() bool {
	return time.Now().After(t.DueDate) && t.Status != "completed"
}
