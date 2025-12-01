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
import type { DeviceType, TicketPriority } from "@/types/ticket";

export const Route = createFileRoute("/tickets/new")({
  component: NewTicket,
});

// Mock technicians data
const mockTechnicians = [
  { id: "1", name: "Sarah Tech" },
  { id: "2", name: "Mike Repair" },
  { id: "3", name: "Lisa Fix" },
];

function NewTicket() {
  const navigate = useNavigate();
  const [deviceType, setDeviceType] = useState<DeviceType | "">("");
  const [priority, setPriority] = useState<TicketPriority>("normal");
  const [assignedTo, setAssignedTo] = useState<string>("");

  const handleSubmit = (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    const formData = new FormData(e.currentTarget);

    // In a real app, this would make an API call
    console.log("Form data:", Object.fromEntries(formData));

    // Redirect to dashboard after creation
    navigate({ to: "/dashboard" });
  };

  return (
    <div className="px-4 py-6 sm:px-0">
      {/* Page Header */}
      <div className="mb-6">
        <h2 className="text-2xl font-bold text-gray-900">Create New Ticket</h2>
        <p className="mt-1 text-sm text-gray-600">
          Fill in the repair ticket details
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
                      placeholder="John Doe"
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
                      placeholder="(555) 123-4567"
                      className="mt-1"
                    />
                  </div>

                  <div>
                    <Label htmlFor="customer_email">Email Address</Label>
                    <Input
                      type="email"
                      name="customer_email"
                      id="customer_email"
                      placeholder="john@example.com"
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
                        <SelectValue placeholder="Select device type" />
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
                      placeholder="Apple, Samsung, etc."
                      className="mt-1"
                    />
                  </div>

                  <div>
                    <Label htmlFor="device_model">Model</Label>
                    <Input
                      type="text"
                      name="device_model"
                      id="device_model"
                      placeholder="iPhone 13, Galaxy S21, etc."
                      className="mt-1"
                    />
                  </div>

                  <div>
                    <Label htmlFor="serial_number">Serial Number / IMEI</Label>
                    <Input
                      type="text"
                      name="serial_number"
                      id="serial_number"
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
                      placeholder="Describe the problem in detail..."
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
                          placeholder="0.00"
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
                        className="mt-1"
                      />
                    </div>
                  </div>

                  <div>
                    <Label htmlFor="internal_notes">Internal Notes</Label>
                    <Textarea
                      name="internal_notes"
                      id="internal_notes"
                      rows={3}
                      placeholder="Notes visible only to staff..."
                      className="mt-1"
                    />
                  </div>
                </div>
              </CardContent>
            </Card>

            {/* Form Actions */}
            <div className="flex justify-end space-x-3">
              <Button
                type="button"
                variant="outline"
                onClick={() => navigate({ to: "/dashboard" })}
              >
                Cancel
              </Button>
              <Button type="submit">Create Ticket</Button>
            </div>
          </form>
        </div>

        {/* Sidebar with helpful info */}
        <div className="lg:col-span-1">
          <Card className="bg-blue-50 border-blue-200">
            <CardContent className="p-4">
              <h4 className="text-sm font-medium text-blue-900 mb-2">
                Quick Tips
              </h4>
              <ul className="text-sm text-blue-700 space-y-2">
                <li>• Collect device password/PIN if needed</li>
                <li>• Document any existing damage</li>
                <li>• Ask about data backup needs</li>
                <li>• Set realistic due dates</li>
                <li>• Take photos if necessary</li>
              </ul>
            </CardContent>
          </Card>
        </div>
      </div>
    </div>
  );
}
