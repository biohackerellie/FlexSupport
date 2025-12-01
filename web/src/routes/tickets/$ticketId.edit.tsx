import { createFileRoute, useNavigate } from "@tanstack/react-router";
import { useState } from "react";
import { Button } from "@/components/ui/button";
import { Card, CardContent } from "@/components/ui/card";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "@/components/ui/select";
import { Textarea } from "@/components/ui/textarea";
import type { Ticket, DeviceType, TicketPriority } from "@/types/ticket";

export const Route = createFileRoute("/tickets/$ticketId/edit")({
  component: EditTicket,
});

// Mock technicians data
const mockTechnicians = [
  { id: "1", name: "Sarah Tech" },
  { id: "2", name: "Mike Repair" },
  { id: "3", name: "Lisa Fix" },
];

// Mock existing ticket data
const mockTicket: Ticket = {
  id: "001",
  customerName: "John Doe",
  customerPhone: "(555) 123-4567",
  customerEmail: "john@example.com",
  deviceType: "smartphone",
  deviceBrand: "Apple",
  deviceModel: "iPhone 13",
  serialNumber: "ABCD1234567890",
  issueDescription: "Screen is cracked and touch is not responding properly.",
  status: "in_progress",
  priority: "high",
  assignedTo: "Sarah Tech",
  dueDate: "2025-11-28",
  estimatedCost: 150,
  totalCost: 180,
  totalPartsCost: 120,
  parts: [],
  notes: [],
  createdAt: "2025-11-24",
  updatedAt: "2025-11-26",
  createdBy: "Front Desk",
};

function EditTicket() {
  const { ticketId } = Route.useParams();
  const navigate = useNavigate();
  const [ticket] = useState<Ticket>(mockTicket);
  const [deviceType, setDeviceType] = useState<DeviceType>(
    ticket.deviceType
  );
  const [priority, setPriority] = useState<TicketPriority>(ticket.priority);
  const [assignedTo, setAssignedTo] = useState<string>(
    ticket.assignedTo || ""
  );

  const handleSubmit = (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    const formData = new FormData(e.currentTarget);

    // In a real app, this would make an API call
    console.log("Form data:", Object.fromEntries(formData));

    // Redirect to ticket view after update
    navigate({ to: "/tickets/$ticketId", params: { ticketId } });
  };

  return (
    <div className="px-4 py-6 sm:px-0">
      {/* Page Header */}
      <div className="mb-6">
        <h2 className="text-2xl font-bold text-gray-900">
          Edit Ticket #{ticket.id}
        </h2>
        <p className="mt-1 text-sm text-gray-600">
          Update the repair ticket details
        </p>
      </div>

      <div className="grid grid-cols-1 lg:grid-cols-3 gap-6">
        {/* Main Form */}
        <div className="lg:col-span-2">
          <form onSubmit={handleSubmit} className="space-y-6">
            {/* Customer Information */}
            <Card>
              <CardContent className="px-4 py-5 sm:p-6">
                <h3 className="text-lg font-medium text-gray-900 mb-4">
                  Customer Information
                </h3>

                <div className="grid grid-cols-1 gap-6 sm:grid-cols-2">
                  <div className="col-span-2">
                    <Label htmlFor="customer_name">
                      Customer Name <span className="text-red-500">*</span>
                    </Label>
                    <Input
                      type="text"
                      name="customer_name"
                      id="customer_name"
                      required
                      defaultValue={ticket.customerName}
                      className="mt-1"
                    />
                  </div>

                  <div>
                    <Label htmlFor="customer_phone">
                      Phone Number <span className="text-red-500">*</span>
                    </Label>
                    <Input
                      type="tel"
                      name="customer_phone"
                      id="customer_phone"
                      required
                      defaultValue={ticket.customerPhone}
                      className="mt-1"
                    />
                  </div>

                  <div>
                    <Label htmlFor="customer_email">Email Address</Label>
                    <Input
                      type="email"
                      name="customer_email"
                      id="customer_email"
                      defaultValue={ticket.customerEmail}
                      className="mt-1"
                    />
                  </div>
                </div>
              </CardContent>
            </Card>

            {/* Device Information */}
            <Card>
              <CardContent className="px-4 py-5 sm:p-6">
                <h3 className="text-lg font-medium text-gray-900 mb-4">
                  Device Information
                </h3>

                <div className="grid grid-cols-1 gap-6 sm:grid-cols-2">
                  <div>
                    <Label htmlFor="device_type">
                      Device Type <span className="text-red-500">*</span>
                    </Label>
                    <Select
                      value={deviceType}
                      onValueChange={(value) =>
                        setDeviceType(value as DeviceType)
                      }
                      required
                    >
                      <SelectTrigger className="mt-1">
                        <SelectValue />
                      </SelectTrigger>
                      <SelectContent>
                        <SelectItem value="smartphone">Smartphone</SelectItem>
                        <SelectItem value="tablet">Tablet</SelectItem>
                        <SelectItem value="laptop">Laptop</SelectItem>
                        <SelectItem value="desktop">Desktop</SelectItem>
                        <SelectItem value="console">Gaming Console</SelectItem>
                        <SelectItem value="other">Other</SelectItem>
                      </SelectContent>
                    </Select>
                    <input
                      type="hidden"
                      name="device_type"
                      value={deviceType}
                    />
                  </div>

                  <div>
                    <Label htmlFor="device_brand">Brand</Label>
                    <Input
                      type="text"
                      name="device_brand"
                      id="device_brand"
                      defaultValue={ticket.deviceBrand}
                      className="mt-1"
                    />
                  </div>

                  <div>
                    <Label htmlFor="device_model">Model</Label>
                    <Input
                      type="text"
                      name="device_model"
                      id="device_model"
                      defaultValue={ticket.deviceModel}
                      className="mt-1"
                    />
                  </div>

                  <div>
                    <Label htmlFor="serial_number">Serial Number / IMEI</Label>
                    <Input
                      type="text"
                      name="serial_number"
                      id="serial_number"
                      defaultValue={ticket.serialNumber}
                      className="mt-1"
                    />
                  </div>
                </div>
              </CardContent>
            </Card>

            {/* Repair Details */}
            <Card>
              <CardContent className="px-4 py-5 sm:p-6">
                <h3 className="text-lg font-medium text-gray-900 mb-4">
                  Repair Details
                </h3>

                <div className="space-y-6">
                  <div>
                    <Label htmlFor="issue_description">
                      Issue Description <span className="text-red-500">*</span>
                    </Label>
                    <Textarea
                      name="issue_description"
                      id="issue_description"
                      required
                      rows={4}
                      defaultValue={ticket.issueDescription}
                      className="mt-1"
                    />
                  </div>

                  <div className="grid grid-cols-1 gap-6 sm:grid-cols-2">
                    <div>
                      <Label htmlFor="priority">Priority</Label>
                      <Select
                        value={priority}
                        onValueChange={(value) =>
                          setPriority(value as TicketPriority)
                        }
                      >
                        <SelectTrigger className="mt-1">
                          <SelectValue />
                        </SelectTrigger>
                        <SelectContent>
                          <SelectItem value="low">Low</SelectItem>
                          <SelectItem value="normal">Normal</SelectItem>
                          <SelectItem value="high">High</SelectItem>
                          <SelectItem value="urgent">Urgent</SelectItem>
                        </SelectContent>
                      </Select>
                      <input type="hidden" name="priority" value={priority} />
                    </div>

                    <div>
                      <Label htmlFor="estimated_cost">Estimated Cost</Label>
                      <div className="mt-1 relative">
                        <span className="absolute inset-y-0 left-3 flex items-center text-gray-500">
                          $
                        </span>
                        <Input
                          type="number"
                          name="estimated_cost"
                          id="estimated_cost"
                          step="0.01"
                          min="0"
                          defaultValue={ticket.estimatedCost}
                          className="pl-7"
                        />
                      </div>
                    </div>

                    <div>
                      <Label htmlFor="assigned_to">Assign To</Label>
                      <Select
                        value={assignedTo}
                        onValueChange={setAssignedTo}
                      >
                        <SelectTrigger className="mt-1">
                          <SelectValue placeholder="Unassigned" />
                        </SelectTrigger>
                        <SelectContent>
                          <SelectItem value="">Unassigned</SelectItem>
                          {mockTechnicians.map((tech) => (
                            <SelectItem key={tech.id} value={tech.id}>
                              {tech.name}
                            </SelectItem>
                          ))}
                        </SelectContent>
                      </Select>
                      <input
                        type="hidden"
                        name="assigned_to"
                        value={assignedTo}
                      />
                    </div>

                    <div>
                      <Label htmlFor="due_date">Due Date</Label>
                      <Input
                        type="date"
                        name="due_date"
                        id="due_date"
                        defaultValue={ticket.dueDate}
                        className="mt-1"
                      />
                    </div>
                  </div>
                </div>
              </CardContent>
            </Card>

            {/* Form Actions */}
            <div className="flex justify-end space-x-3">
              <Button
                type="button"
                variant="outline"
                onClick={() =>
                  navigate({ to: "/tickets/$ticketId", params: { ticketId } })
                }
              >
                Cancel
              </Button>
              <Button type="submit">Update Ticket</Button>
            </div>
          </form>
        </div>

        {/* Sidebar with ticket history */}
        <div className="lg:col-span-1">
          <Card>
            <CardContent className="p-4">
              <h4 className="text-sm font-medium text-gray-900 mb-3">
                Ticket History
              </h4>
              <div className="space-y-3">
                <div className="text-xs">
                  <span className="text-gray-500">Created:</span>
                  <span className="text-gray-900 ml-1">
                    {ticket.createdAt}
                  </span>
                </div>
                {ticket.updatedAt && (
                  <div className="text-xs">
                    <span className="text-gray-500">Last Updated:</span>
                    <span className="text-gray-900 ml-1">
                      {ticket.updatedAt}
                    </span>
                  </div>
                )}
                {ticket.createdBy && (
                  <div className="text-xs">
                    <span className="text-gray-500">Created By:</span>
                    <span className="text-gray-900 ml-1">
                      {ticket.createdBy}
                    </span>
                  </div>
                )}
              </div>
            </CardContent>
          </Card>
        </div>
      </div>
    </div>
  );
}
