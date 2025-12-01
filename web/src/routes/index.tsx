import { createFileRoute, redirect } from "@tanstack/react-router";

export const Route = createFileRoute("/")({
  beforeLoad: () => {
    // Redirect to dashboard as the default page
    throw redirect({
      to: "/dashboard",
    });
  },
});
