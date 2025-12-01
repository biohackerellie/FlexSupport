import { createFileRoute, Link } from "@tanstack/react-router";
import { useState } from "react";
import {
  Clock,
  ShoppingCart,
  CheckCircle,
  Trash2,
  AlertCircle,
} from "lucide-react";
import { Button } from "@/components/ui/button";
import { Card, CardContent } from "@/components/ui/card";
import { Input } from "@/components/ui/input";
import { Textarea } from "@/components/ui/textarea";
import type { Ticket, Part, Note } from "@/types/ticket";
import {
  getStatusBadgeClass,
  getStatusDisplayText,
} from "@/types/ticket";

export const Route = createFileRoute("/tickets/$ticketId")({
  component: TechnicianView,
});

// Mock ticket data
const mockTicket: Ticket = {
  id: "001",
  customerName: "John Doe",
  customerPhone: "(555) 123-4567",
  customerEmail: "john@example.com",
  deviceType: "smartphone",
  deviceBrand: "Apple",
  deviceModel: "iPhone 13",
  serialNumber: "ABCD1234567890",
  issueDescription:
    "Screen is cracked and touch is not responding properly. Customer reports the phone was dropped from about 3 feet height.",
  status: "in_progress",
  priority: "high",
  assignedTo: "Sarah Tech",
  dueDate: "2025-11-28",
  estimatedCost: 150,
  totalCost: 180,
  totalPartsCost: 120,
  parts: [
    {
      id: "p1",
      name: "iPhone 13 Screen Assembly",
      quantity: 1,
      cost: 120,
    },
  ],
  notes: [
    {
      id: "n1",
      author: "Sarah Tech",
      content: "Inspected device. Screen replacement required.",
      timestamp: "2025-11-26 10:30 AM",
    },
  ],
  createdAt: "2025-11-24",
  updatedAt: "2025-11-26",
  createdBy: "Front Desk",
};

function TechnicianView() {
  const { ticketId } = Route.useParams();
  const [ticket, setTicket] = useState<Ticket>(mockTicket);
  const [showAddPart, setShowAddPart] = useState(false);
  const [newNote, setNewNote] = useState("");

  const handleStatusUpdate = (newStatus: Ticket["status"]) => {
    setTicket((prev) => ({ ...prev, status: newStatus }));
  };

  const handleAddPart = (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    const formData = new FormData(e.currentTarget);
    const newPart: Part = {
      id: `p${ticket.parts.length + 1}`,
      name: formData.get("part_name") as string,
      quantity: Number(formData.get("quantity")),
      cost: Number(formData.get("cost")),
    };
    setTicket((prev) => ({
      ...prev,
      parts: [...prev.parts, newPart],
      totalPartsCost: prev.totalPartsCost + newPart.cost * newPart.quantity,
    }));
    setShowAddPart(false);
    e.currentTarget.reset();
  };

  const handleDeletePart = (partId: string) => {
    setTicket((prev) => {
      const deletedPart = prev.parts.find((p) => p.id === partId);
      const costReduction = deletedPart
        ? deletedPart.cost * deletedPart.quantity
        : 0;
      return {
        ...prev,
        parts: prev.parts.filter((p) => p.id !== partId),
        totalPartsCost: prev.totalPartsCost - costReduction,
      };
    });
  };

  const handleAddNote = (e: React.FormEvent) => {
    e.preventDefault();
    if (!newNote.trim()) return;

    const note: Note = {
      id: `n${ticket.notes.length + 1}`,
      author: "Current User",
      content: newNote,
      timestamp: new Date().toLocaleString(),
    };
    setTicket((prev) => ({
      ...prev,
      notes: [note, ...prev.notes],
    }));
    setNewNote("");
  };

  return (
    <div className="px-4 py-6 sm:px-0">
      {/* Page Header */}
      <div className="mb-6 flex justify-between items-start">
        <div>
          <div className="flex items-center gap-3">
            <h2 className="text-2xl font-bold text-gray-900">
              Ticket #{ticket.id}
            </h2>
            <span
              className={`px-3 py-1 inline-flex text-sm leading-5 font-semibold rounded-full ${getStatusBadgeClass(ticket.status)}`}
            >
              {getStatusDisplayText(ticket.status)}
            </span>
          </div>
          <p className="mt-1 text-sm text-gray-600 capitalize">
            {ticket.deviceType} - {ticket.deviceBrand} {ticket.deviceModel}
          </p>
        </div>
        <Link
          to="/technician"
          className="text-sm text-gray-600 hover:text-gray-900"
        >
          ← Back to queue
        </Link>
      </div>

      <div className="grid grid-cols-1 lg:grid-cols-3 gap-6">
        {/* Main Content */}
        <div className="lg:col-span-2 space-y-6">
          {/* Quick Status Update */}
          <Card>
            <CardContent className="px-4 py-5 sm:p-6">
              <h3 className="text-lg font-medium text-gray-900 mb-4">
                Quick Actions
              </h3>

              <div className="flex flex-wrap gap-2">
                <Button
                  variant="outline"
                  onClick={() => handleStatusUpdate("in_progress")}
                >
                  <Clock className="h-4 w-4 mr-1.5 text-yellow-500" />
                  Start Work
                </Button>

                <Button
                  variant="outline"
                  onClick={() => handleStatusUpdate("waiting_parts")}
                >
                  <ShoppingCart className="h-4 w-4 mr-1.5 text-orange-500" />
                  Waiting for Parts
                </Button>

                <Button
                  variant="outline"
                  onClick={() => handleStatusUpdate("ready")}
                >
                  <CheckCircle className="h-4 w-4 mr-1.5 text-blue-500" />
                  Ready for Pickup
                </Button>

                <Button
                  variant="default"
                  className="bg-green-600 hover:bg-green-700"
                  onClick={() => {
                    if (
                      confirm(
                        "Are you sure you want to mark this ticket as completed?"
                      )
                    ) {
                      handleStatusUpdate("completed");
                    }
                  }}
                >
                  <CheckCircle className="h-4 w-4 mr-1.5" />
                  Mark Completed
                </Button>
              </div>
            </CardContent>
          </Card>

          {/* Issue Details */}
          <Card>
            <CardContent className="px-4 py-5 sm:p-6">
              <h3 className="text-lg font-medium text-gray-900 mb-4">
                Issue Description
              </h3>
              <p className="text-gray-700 whitespace-pre-line">
                {ticket.issueDescription}
              </p>
            </CardContent>
          </Card>

          {/* Parts & Materials */}
          <Card>
            <CardContent className="px-4 py-5 sm:p-6">
              <div className="flex justify-between items-center mb-4">
                <h3 className="text-lg font-medium text-gray-900">
                  Parts & Materials
                </h3>
                <Button
                  variant="link"
                  onClick={() => setShowAddPart(!showAddPart)}
                >
                  + Add Part
                </Button>
              </div>

              {/* Add Part Form */}
              {showAddPart && (
                <div className="mb-4 p-4 bg-gray-50 rounded-md">
                  <form onSubmit={handleAddPart}>
                    <div className="grid grid-cols-1 gap-4 sm:grid-cols-4">
                      <div className="sm:col-span-2">
                        <Input
                          type="text"
                          name="part_name"
                          placeholder="Part name"
                          required
                        />
                      </div>
                      <div>
                        <Input
                          type="number"
                          name="quantity"
                          placeholder="Qty"
                          min="1"
                          defaultValue="1"
                          required
                        />
                      </div>
                      <div>
                        <div className="relative">
                          <span className="absolute inset-y-0 left-3 flex items-center text-gray-500">
                            $
                          </span>
                          <Input
                            type="number"
                            name="cost"
                            placeholder="Cost"
                            step="0.01"
                            min="0"
                            className="pl-7"
                            required
                          />
                        </div>
                      </div>
                    </div>
                    <div className="mt-2 flex justify-end gap-2">
                      <Button
                        type="button"
                        variant="ghost"
                        onClick={() => setShowAddPart(false)}
                      >
                        Cancel
                      </Button>
                      <Button type="submit">Add Part</Button>
                    </div>
                  </form>
                </div>
              )}

              {/* Parts List */}
              <div className="space-y-2">
                {ticket.parts.length > 0 ? (
                  ticket.parts.map((part) => (
                    <div
                      key={part.id}
                      className="flex items-center justify-between p-3 bg-gray-50 rounded-md"
                    >
                      <div className="flex-1">
                        <span className="text-sm font-medium text-gray-900">
                          {part.name}
                        </span>
                        <span className="text-sm text-gray-500 ml-2">
                          × {part.quantity}
                        </span>
                      </div>
                      <div className="flex items-center gap-3">
                        <span className="text-sm font-medium text-gray-900">
                          ${part.cost.toFixed(2)}
                        </span>
                        <Button
                          variant="ghost"
                          size="icon"
                          className="text-red-600 hover:text-red-900 h-8 w-8"
                          onClick={() => handleDeletePart(part.id)}
                        >
                          <Trash2 className="h-4 w-4" />
                        </Button>
                      </div>
                    </div>
                  ))
                ) : (
                  <p className="text-sm text-gray-500 text-center py-4">
                    No parts added yet
                  </p>
                )}
              </div>

              {ticket.parts.length > 0 && (
                <div className="mt-4 pt-4 border-t border-gray-200">
                  <div className="flex justify-between text-sm">
                    <span className="font-medium text-gray-900">
                      Total Parts Cost:
                    </span>
                    <span className="font-bold text-gray-900">
                      ${ticket.totalPartsCost.toFixed(2)}
                    </span>
                  </div>
                </div>
              )}
            </CardContent>
          </Card>

          {/* Work Log / Notes */}
          <Card>
            <CardContent className="px-4 py-5 sm:p-6">
              <h3 className="text-lg font-medium text-gray-900 mb-4">
                Work Log
              </h3>

              {/* Add Note Form */}
              <form onSubmit={handleAddNote} className="mb-4">
                <Textarea
                  value={newNote}
                  onChange={(e) => setNewNote(e.target.value)}
                  rows={3}
                  placeholder="Add a work note..."
                  required
                />
                <div className="mt-2 flex justify-end">
                  <Button type="submit">Add Note</Button>
                </div>
              </form>

              {/* Notes List */}
              <div className="space-y-3">
                {ticket.notes.length > 0 ? (
                  ticket.notes.map((note) => (
                    <div key={note.id} className="p-3 bg-gray-50 rounded-md">
                      <div className="flex justify-between items-start mb-1">
                        <span className="text-sm font-medium text-gray-900">
                          {note.author}
                        </span>
                        <span className="text-xs text-gray-500">
                          {note.timestamp}
                        </span>
                      </div>
                      <p className="text-sm text-gray-700 whitespace-pre-line">
                        {note.content}
                      </p>
                    </div>
                  ))
                ) : (
                  <p className="text-sm text-gray-500 text-center py-4">
                    No work notes yet
                  </p>
                )}
              </div>
            </CardContent>
          </Card>
        </div>

        {/* Sidebar */}
        <div className="lg:col-span-1 space-y-6">
          {/* Customer Info */}
          <Card>
            <CardContent className="px-4 py-5 sm:p-6">
              <h3 className="text-sm font-medium text-gray-900 mb-3">
                Customer Information
              </h3>
              <dl className="space-y-3">
                <div>
                  <dt className="text-xs text-gray-500">Name</dt>
                  <dd className="text-sm font-medium text-gray-900">
                    {ticket.customerName}
                  </dd>
                </div>
                <div>
                  <dt className="text-xs text-gray-500">Phone</dt>
                  <dd className="text-sm text-gray-900">
                    <a
                      href={`tel:${ticket.customerPhone}`}
                      className="text-blue-600 hover:text-blue-900"
                    >
                      {ticket.customerPhone}
                    </a>
                  </dd>
                </div>
                {ticket.customerEmail && (
                  <div>
                    <dt className="text-xs text-gray-500">Email</dt>
                    <dd className="text-sm text-gray-900">
                      <a
                        href={`mailto:${ticket.customerEmail}`}
                        className="text-blue-600 hover:text-blue-900"
                      >
                        {ticket.customerEmail}
                      </a>
                    </dd>
                  </div>
                )}
              </dl>
            </CardContent>
          </Card>

          {/* Device Details */}
          <Card>
            <CardContent className="px-4 py-5 sm:p-6">
              <h3 className="text-sm font-medium text-gray-900 mb-3">
                Device Details
              </h3>
              <dl className="space-y-3">
                <div>
                  <dt className="text-xs text-gray-500">Type</dt>
                  <dd className="text-sm text-gray-900 capitalize">
                    {ticket.deviceType}
                  </dd>
                </div>
                <div>
                  <dt className="text-xs text-gray-500">Brand/Model</dt>
                  <dd className="text-sm text-gray-900">
                    {ticket.deviceBrand} {ticket.deviceModel}
                  </dd>
                </div>
                {ticket.serialNumber && (
                  <div>
                    <dt className="text-xs text-gray-500">Serial/IMEI</dt>
                    <dd className="text-sm text-gray-900 font-mono">
                      {ticket.serialNumber}
                    </dd>
                  </div>
                )}
              </dl>
            </CardContent>
          </Card>

          {/* Ticket Metadata */}
          <Card>
            <CardContent className="px-4 py-5 sm:p-6">
              <h3 className="text-sm font-medium text-gray-900 mb-3">
                Ticket Details
              </h3>
              <dl className="space-y-3">
                <div>
                  <dt className="text-xs text-gray-500">Priority</dt>
                  <dd className="text-sm text-gray-900 capitalize">
                    {ticket.priority}
                  </dd>
                </div>
                <div>
                  <dt className="text-xs text-gray-500">Created</dt>
                  <dd className="text-sm text-gray-900">
                    {ticket.createdAt}
                  </dd>
                </div>
                {ticket.dueDate && (
                  <div>
                    <dt className="text-xs text-gray-500">Due Date</dt>
                    <dd
                      className={`text-sm text-gray-900 ${ticket.isOverdue ? "text-red-600 font-semibold" : ""}`}
                    >
                      {ticket.dueDate}
                      {ticket.isOverdue && (
                        <AlertCircle className="inline h-4 w-4 ml-1" />
                      )}
                    </dd>
                  </div>
                )}
                <div>
                  <dt className="text-xs text-gray-500">Estimated Cost</dt>
                  <dd className="text-sm text-gray-900">
                    ${ticket.estimatedCost?.toFixed(2)}
                  </dd>
                </div>
                <div>
                  <dt className="text-xs text-gray-500">Total Cost</dt>
                  <dd className="text-lg font-bold text-gray-900">
                    ${ticket.totalCost.toFixed(2)}
                  </dd>
                </div>
              </dl>
            </CardContent>
          </Card>

          {/* Actions */}
          <Card>
            <CardContent className="px-4 py-5 sm:p-6">
              <Button variant="outline" className="w-full" asChild>
                <Link
                  to="/tickets/$ticketId/edit"
                  params={{ ticketId: ticket.id }}
                >
                  Edit Ticket
                </Link>
              </Button>
            </CardContent>
          </Card>
        </div>
      </div>
    </div>
  );
}
