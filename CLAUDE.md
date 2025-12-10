# FlexSupport - Project Documentation for Claude

## Project Overview

FlexSupport is a lightweight helpdesk and ticketing system specifically tailored for **repair shop workflows**. It's designed to manage repair tickets for items like boots, shoes, bags, and other goods, tracking customer information, repair status, parts, and work notes.

**Primary Use Case**: Repair shops that need to track customer repairs from intake through completion, including parts management and technician assignment.

## Technology Stack

### Backend
- **Go 1.25** - Primary language
- **Chi (v5)** - HTTP router and middleware
- **Templ** - Type-safe Go templating (generates `*_templ.go` files)
- **Shopify Integration** - Customer/order integration via `go-shopify` SDK

### Frontend
- **htmx** - Dynamic HTML interactions without heavy JavaScript
- **Alpine.js** - Lightweight JavaScript framework for interactivity
- **Tailwind CSS** - Utility-first CSS framework
- **TemplUI** - Custom component library (see `ui/components/`)

### Development Tools
- **Air** - Hot reloading for Go server (`.air.toml`)
- **Task** - Task runner (Taskfile.yml) for build automation
- **godotenv** - Environment variable management

## Project Structure

```
FlexSupport/
├── main.go              # Entry point, runs App()
├── app.go               # Main application logic and server setup
├── .env                 # Environment variables (not committed)
├── go.mod               # Go dependencies
├── Taskfile.yml         # Task definitions for dev workflow
├── .air.toml            # Air hot reload configuration
│
├── cmd/                 # Command-line tools (if any)
│
├── internal/
│   ├── config/          # Application configuration
│   │   └── config.go    # Config struct and env var loading
│   │
│   ├── models/          # Data models
│   │   └── ticket.go    # Ticket, Part, WorkNote, Customer, Technician, TicketStats
│   │
│   ├── handlers/        # HTTP request handlers (stub/interface level)
│   │   ├── handlers.go  # Handler struct with methods for endpoints
│   │   └── getHome.go   # Home page handler
│   │
│   ├── routes/          # Feature-specific route handlers
│   │   ├── dashboard/   # Dashboard page and logic
│   │   │   ├── dashboard.templ       # Templ template
│   │   │   ├── dashboard_templ.go    # Generated Go code
│   │   │   └── handler.go
│   │   └── tickets/     # Ticket management routes
│   │       ├── tickets.go
│   │       ├── ticket-form.templ
│   │       └── ticket-page.templ
│   │
│   ├── router/          # Route registration and middleware setup
│   │   └── router.go    # Chi router configuration
│   │
│   ├── server/          # HTTP server configuration
│   │   └── server.go
│   │
│   ├── middleware/      # Custom middleware
│   │   ├── middleware.go
│   │   ├── logging.go   # Request logging
│   │   └── noonce.go    # Security nonce handling
│   │
│   ├── layout/          # Page layouts and common UI
│   │   ├── base.templ   # Base HTML layout
│   │   ├── base_templ.go
│   │   └── layout.go
│   │
│   ├── lib/             # Utility libraries
│   │   └── logger/      # Structured logging setup
│   │
│   ├── shopify/         # Shopify integration code
│   │   └── shopify.go
│   │
│   ├── utils/           # General utilities
│   │   └── templui.go   # TemplUI component helpers
│   │
│   └── ports/           # Interface definitions (ports pattern)
│       └── ports.go
│
├── static/              # Static assets
│   ├── assets/
│   │   ├── css/
│   │   │   ├── input.css      # Tailwind source
│   │   │   └── output.min.css # Compiled CSS
│   │   ├── js/          # JavaScript files
│   │   ├── icons/       # Icon assets
│   │   └── public/      # Public assets
│   └── assets.go        # Go embed for static files
│
├── ui/                  # Templ components
│   ├── components/      # Reusable UI components
│   │   ├── button/
│   │   ├── card/
│   │   ├── form/
│   │   ├── input/
│   │   ├── table/
│   │   └── ... (many more)
│   └── modules/         # Larger UI modules
│       └── navbar.templ
│
└── tmp/                 # Build artifacts (Air output)
    └── main             # Compiled binary
```

## Core Data Models

### Ticket
The central entity representing a repair job. Key fields:
- **Status**: `new`, `in_progress`, `waiting_parts`, `ready`, `completed`
- **Priority**: `low`, `normal`, `high`, `urgent`
- **Customer Info**: Name, phone, email
- **Item Info**: Type (boot/shoe/bag/other), brand, model, serial number
- **Repair Details**: Issue description, internal notes, estimated cost
- **Assignment**: Assigned technician, due date
- **Relations**: Parts (replacement parts used), WorkNotes (log entries)

### Part
Replacement parts/materials used in repairs (quantity, cost)

### WorkNote
Log entries and notes added to tickets by technicians

### Customer & Technician
User entities (planned for future expansion)

## Routing Structure

**Main Router**: `internal/router/router.go`

### Routes
- `GET /` - Dashboard (home page)
- `GET /assets/*` - Static assets
- `POST /tickets` - Create new ticket
- `GET /tickets/search` - Search tickets
- `POST /tickets/{id}` - Update ticket
- `POST /tickets/{id}/status` - Update ticket status
- `POST /tickets/{id}/parts` - Add part to ticket
- `DELETE /tickets/{id}/parts/{partId}` - Remove part
- `POST /tickets/{id}/notes` - Add work note
- `GET /api/stats/open` - Get open ticket count (API endpoint for htmx)

## Development Workflow

### Starting Development

```bash
# Option 1: Full development mode (recommended)
task gen    # Generates Templ files + Tailwind CSS
task dev    # Starts Air hot-reload server

# Option 2: Individual tasks
task templ     # Generate Templ files only
task tailwind  # Build Tailwind CSS only
go run .       # Run server directly
```

**Server runs on**: `http://localhost:8080`

### File Generation

**Templ Files**: `*.templ` → `*_templ.go` (auto-generated, DON'T edit)
- Templ is a templating language that compiles to Go code
- Edit `.templ` files, never edit `_templ.go` files
- Run `task templ` to regenerate

**Tailwind CSS**: `static/assets/css/input.css` → `output.min.css`
- Run `task tailwind` to rebuild CSS
- Minified output for production

### Hot Reloading

**Air** watches for changes to:
- `*.go` files (except `*_templ.go`)
- `*.templ` files
- `*.html` files

**Important**: Air excludes `_templ.go` from watch list to prevent rebuild loops. You must manually run `task templ` when editing `.templ` files, or use `task gen` which runs both templ and tailwind in watch mode.

## Key Patterns and Conventions

### 1. Handler Pattern
- `internal/handlers/handlers.go` - Handler struct with all endpoint methods
- Route-specific handlers in `internal/routes/*/handler.go`
- Handlers receive `*slog.Logger` for structured logging

### 2. Templ Components
- Type-safe Go templating
- Components are Go functions returning `templ.Component`
- Call components from other components or HTTP handlers
- Example: `dashboard.Show()` returns a component to render

### 3. htmx Integration
- Partial page updates without full page reloads
- Server returns HTML fragments
- Use `hx-*` attributes in Templ templates

### 4. Middleware Stack
- Logger, Recoverer (panic recovery)
- Custom logging middleware
- Content-Type enforcement (`text/html`)
- CSP (Content Security Policy) - currently commented out

### 5. Environment Configuration
- `development` vs `production` mode
- Config loaded from environment variables via `.env`
- Config struct in `internal/config/config.go`

## Common Tasks

### Adding a New Route
1. Add handler method to `internal/handlers/handlers.go`
2. Register route in `internal/router/router.go`
3. Create Templ template if needed in `internal/routes/*/`
4. Run `task templ` to generate Go code

### Creating a New Templ Component
1. Create `component.templ` in appropriate directory
2. Write templ code using `templ` syntax
3. Run `task templ` to generate `component_templ.go`
4. Import and use in handlers or other components

### Adding Tailwind Classes
1. Use classes in Templ files
2. Tailwind will auto-detect and include them
3. Run `task tailwind` to rebuild CSS

### Working with Models
- Models are in `internal/models/`
- Use struct tags for JSON/DB serialization
- Add helper methods for display logic (e.g., `StatusClass()`, `StatusDisplay()`)

## Important Files

- **`main.go`**: Entry point, loads env vars, calls `App()`
- **`app.go`**: Creates config, logger, handlers, router, starts HTTP server
- **`internal/router/router.go`**: All route definitions
- **`internal/models/ticket.go`**: Core data structures
- **`.air.toml`**: Hot reload configuration (watch patterns, build commands)
- **`Taskfile.yml`**: Development tasks (templ, tailwind, dev server)
- **`.env`**: Environment variables (LOG_LEVEL, ENVIRONMENT, etc.)

## Environment Variables

Default values in `internal/config/config.go`:
- `LOG_LEVEL`: `debug` (development), `info` (production)
- `VERBOSE_LOGGING`: `false`
- `ENVIRONMENT`: `development` | `production`
- `DOMAIN`: `http://localhost:8080`

## Database / Persistence

**Currently**: No database implementation visible in codebase
**Likely**: Using in-memory data or planning to add database layer via `internal/ports/` (ports/adapters pattern)

If implementing database:
- Add repository interfaces in `internal/ports/`
- Create implementations (e.g., PostgreSQL, SQLite)
- Inject into handlers via dependency injection

## Tips for Working with This Codebase

1. **Always regenerate Templ files** after editing `.templ` files
2. **Don't edit `_templ.go` files** - they're auto-generated
3. **Use `task gen`** for concurrent Templ + Tailwind watching during development
4. **Check Air logs** in `tmp/build-errors.log` if builds fail
5. **Use structured logging** via `slog.Logger` passed to handlers
6. **Follow Chi router patterns** - use `r.Group()` for middleware scoping
7. **Templ components are type-safe** - leverage Go's type system
8. **htmx expects HTML fragments** - return partial templates for AJAX endpoints
9. **Static assets are embedded** via `static/assets.go` (Go 1.16+ embed)
10. **Middleware order matters** - logger before recoverer, auth after logging

## Testing Notes

- No test files currently visible in the repository
- Consider adding tests in `*_test.go` files alongside code
- Integration tests for handlers
- Unit tests for models and utilities

## Future Enhancements / TODOs

Based on commented-out code in `router.go`:
- Ticket list view (`GET /tickets`)
- New ticket form (`GET /tickets/new`)
- View single ticket (`GET /tickets/{id}`)
- Edit ticket form (`GET /tickets/{id}/edit`)
- Content Security Policy middleware (currently disabled)
- Right now the system is focused on a cobler repair shop workflow as they are the first clients I'm working with. The end goal is to make this more generic and adaptable to different  workflows such as car repair or an IT team ticket system.

## Shopify Integration

- Located in `internal/shopify/shopify.go`
- Uses `github.com/bold-commerce/go-shopify/v4`
- Likely for importing customer/order data
- Integration details to be explored

## Quick Reference

**Start dev server**: `task gen && task dev`
**Build only**: `go build -o ./tmp/main .`
**Generate Templ**: `task templ` or `go tool templ generate`
**Build CSS**: `task tailwind`
**Server port**: `:8080`
**Logs**: Structured JSON (prod) or text (dev) via `slog`

---

**Last Updated**: 2025-12-10
**Go Version**: 1.25
**Main Branch**: `main`
