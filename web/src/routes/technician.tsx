import { createFileRoute, Link } from "@tanstack/react-router";
import { Card, CardContent } from "@/components/ui/card";
import type { Ticket } from "@/types/ticket";
import { getStatusBadgeClass, getStatusDisplayText } from "@/types/ticket";

export const Route = createFileRoute("/technician")({
  component: TechnicianQueue,
});

// Mock tickets for the technician queue
const mockTickets: Ticket[] = [
  {
    id: "003",
    customerName: "Bob Smith",
    customerPhone: "(555) 456-7890",
    itemType: "hat",
    itemBrand: "Ray-Ban",
    itemModel: "Glasses",
    issueDescription: "needs dry cleaning",
    status: "waiting_parts",
    priority: "urgent",
    assignedTo: "Austin Dipshit",
    dueDate: "2025-11-27",
    estimatedCost: 100,
    totalCost: 100,
    totalPartsCost: 80,
    parts: [],
    notes: [],
    createdAt: "2025-11-25",
    isOverdue: true,
  },
  {
    id: "002",
    customerName: "Jane Smith",
    customerPhone: "(555) 987-6543",
    itemType: "boot",
    itemBrand: "Keen",
    itemModel: "Big Boy Boot",
    issueDescription: "its fucked up",
    status: "new",
    priority: "normal",
    estimatedCost: 200,
    totalCost: 0,
    totalPartsCost: 0,
    parts: [],
    notes: [],
    createdAt: "2025-11-26",
  },
];

function TechnicianQueue() {
  // Group tickets by status for better organization
  const inProgressTickets = mockTickets.filter(
    (t) => t.status === "in_progress",
  );
  const waitingPartsTickets = mockTickets.filter(
    (t) => t.status === "waiting_parts",
  );
  const newTickets = mockTickets.filter((t) => t.status === "new");
  const readyTickets = mockTickets.filter((t) => t.status === "ready");

  return (
    <div className="px-4 py-6 sm:px-0">
      {/* Page Header */}
      <div className="mb-6">
        <h2 className="text-2xl font-bold text-gray-900">
          Technician Work Queue
        </h2>
        <p className="mt-1 text-sm text-gray-600">
          View and manage active repair jobs
        </p>
      </div>

      <div className="space-y-6">
        {/* In Progress Section */}
        <div>
          <h3 className="text-lg font-semibold text-gray-900 mb-3">
            In Progress ({inProgressTickets.length})
          </h3>
          <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
            {inProgressTickets.length > 0 ? (
              inProgressTickets.map((ticket) => (
                <TicketCard key={ticket.id} ticket={ticket} />
              ))
            ) : (
              <p className="text-sm text-gray-500 col-span-full">
                No tickets in progress
              </p>
            )}
          </div>
        </div>

        {/* Waiting for Parts Section */}
        <div>
          <h3 className="text-lg font-semibold text-gray-900 mb-3">
            Waiting for Parts ({waitingPartsTickets.length})
          </h3>
          <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
            {waitingPartsTickets.length > 0 ? (
              waitingPartsTickets.map((ticket) => (
                <TicketCard key={ticket.id} ticket={ticket} />
              ))
            ) : (
              <p className="text-sm text-gray-500 col-span-full">
                No tickets waiting for parts
              </p>
            )}
          </div>
        </div>

        {/* New/Unassigned Section */}
        <div>
          <h3 className="text-lg font-semibold text-gray-900 mb-3">
            New Tickets ({newTickets.length})
          </h3>
          <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
            {newTickets.length > 0 ? (
              newTickets.map((ticket) => (
                <TicketCard key={ticket.id} ticket={ticket} />
              ))
            ) : (
              <p className="text-sm text-gray-500 col-span-full">
                No new tickets
              </p>
            )}
          </div>
        </div>

        {/* Ready for Pickup Section */}
        <div>
          <h3 className="text-lg font-semibold text-gray-900 mb-3">
            Ready for Pickup ({readyTickets.length})
          </h3>
          <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
            {readyTickets.length > 0 ? (
              readyTickets.map((ticket) => (
                <TicketCard key={ticket.id} ticket={ticket} />
              ))
            ) : (
              <p className="text-sm text-gray-500 col-span-full">
                No tickets ready for pickup
              </p>
            )}
          </div>
        </div>
      </div>
    </div>
  );
}

function TicketCard({ ticket }: { ticket: Ticket }) {
  return (
    <Link
      to="/tickets/$ticketId"
      params={{ ticketId: ticket.id }}
      className="block"
    >
      <Card className="hover:shadow-lg transition-shadow cursor-pointer">
        <CardContent className="p-4">
          <div className="flex justify-between items-start mb-2">
            <div>
              <h4 className="text-sm font-semibold text-gray-900">
                #{ticket.id}
              </h4>
              <p className="text-xs text-gray-500">{ticket.customerName}</p>
            </div>
            <span
              className={`px-2 py-1 text-xs font-semibold rounded-full ${getStatusBadgeClass(ticket.status)}`}
            >
              {getStatusDisplayText(ticket.status)}
            </span>
          </div>

          <div className="space-y-1 mb-3">
            <p className="text-sm font-medium text-gray-700 capitalize">
              {ticket.itemType}
            </p>
            {ticket.itemBrand && ticket.itemModel && (
              <p className="text-xs text-gray-500">
                {ticket.itemBrand} {ticket.itemModel}
              </p>
            )}
          </div>

          <p className="text-xs text-gray-600 line-clamp-2 mb-3">
            {ticket.issueDescription}
          </p>

          <div className="flex justify-between items-center text-xs">
            <div>
              {ticket.priority && (
                <span
                  className={`capitalize ${ticket.priority === "urgent" || ticket.priority === "high"
                      ? "text-red-600 font-semibold"
                      : "text-gray-600"
                    }`}
                >
                  {ticket.priority} priority
                </span>
              )}
            </div>
            {ticket.dueDate && (
              <span
                className={`${ticket.isOverdue ? "text-red-600 font-semibold" : "text-gray-500"}`}
              >
                Due: {ticket.dueDate}
              </span>
            )}
          </div>

          {ticket.assignedTo && (
            <div className="mt-2 pt-2 border-t border-gray-200">
              <p className="text-xs text-gray-500">
                Assigned to: {ticket.assignedTo}
              </p>
            </div>
          )}
        </CardContent>
      </Card>
    </Link>
  );
}
