export type TicketStatus =
  | "new"
  | "in_progress"
  | "waiting_parts"
  | "ready"
  | "completed";

export type TicketPriority = "low" | "normal" | "high" | "urgent";

export type ItemType = "boot" | "shoe" | "bag" | "hat" | "other";

export interface Part {
  id: string;
  name: string;
  quantity: number;
  cost: number;
}

export interface Note {
  id: string;
  author: string;
  content: string;
  timestamp: string;
}

export interface Ticket {
  id: string;
  customerName: string;
  customerPhone: string;
  customerEmail?: string;
  itemType: ItemType;
  itemBrand?: string;
  itemModel?: string;
  serialNumber?: string;
  issueDescription: string;
  status: TicketStatus;
  priority: TicketPriority;
  assignedTo?: string;
  dueDate?: string;
  estimatedCost?: number;
  totalCost: number;
  totalPartsCost: number;
  parts: Part[];
  notes: Note[];
  createdAt: string;
  updatedAt?: string;
  createdBy?: string;
  isOverdue?: boolean;
}

export interface DashboardStats {
  openTickets: number;
  inProgress: number;
  overdue: number;
  completedToday: number;
}

export interface Technician {
  id: string;
  name: string;
}

// Helper function to get status badge classes
export function getStatusBadgeClass(status: TicketStatus): string {
  const classes = {
    new: "bg-blue-100 text-blue-800",
    in_progress: "bg-yellow-100 text-yellow-800",
    waiting_parts: "bg-orange-100 text-orange-800",
    ready: "bg-green-100 text-green-800",
    completed: "bg-gray-100 text-gray-800",
  };
  return classes[status];
}

// Helper function to get status display text
export function getStatusDisplayText(status: TicketStatus): string {
  const displayText = {
    new: "New",
    in_progress: "In Progress",
    waiting_parts: "Waiting for Parts",
    ready: "Ready for Pickup",
    completed: "Completed",
  };
  return displayText[status];
}
