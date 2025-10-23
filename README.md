# dfcsr - Dog Registry Web Application

A simple Go web application demonstrating clean architecture patterns with a dog registry system. The application provides HTTP endpoints to view and filter dog information.

## Features

- View all dogs with sortable columns (name, breed)
- Search for individual dogs by name
- Clean layered architecture (Controller → Service → Repository)
- Server-side HTML rendering with Go templates
- In-memory data storage

## Project Structure

```
.
├── cmd/web/          # Application entry point
│   └── main.go       # HTTP server setup and routing
├── internal/dog/     # Dog domain logic
│   ├── controller.go # HTTP handlers and request/response handling
│   ├── service.go    # Business logic layer
│   └── repo_memory.go # In-memory data repository
├── templates/        # HTML templates
│   └── templates.tmpl # View templates for rendering
├── Makefile          # Build and development commands
└── .air.toml         # Hot reload configuration
```

## Architecture

The application follows a three-tier architecture:

1. **Controller Layer** (`controller.go`) - Handles HTTP requests, form parsing, and template rendering
2. **Service Layer** (`service.go`) - Contains business logic, validation, and data transformation
3. **Repository Layer** (`repo_memory.go`) - Manages data persistence (currently in-memory)

## API Endpoints

- `GET /dog/all?order=<name|breed>&direction=<asc|desc>` - List all dogs with optional sorting
- `GET /dog?name=<name>` - Get a specific dog by name

## Dependencies

- [chi/v5](https://github.com/go-chi/chi) - Lightweight HTTP router
- Go 1.25.3

## Development

### Prerequisites
- Go 1.25.3 or later
- [Air](https://github.com/cosmtrek/air) for hot reloading (optional)

### Available Commands

```bash
make help   # Show available commands
make web    # Run web server with hot reload (requires Air)
make lint   # Run linter
```

### Running the Application

Start the server:
```bash
make web
# or
air
```

The server will listen on `http://localhost:8080`
