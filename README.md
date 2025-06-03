# Orb

**Orb** is a small, extensible Go application designed for performing infrastructure connectivity checks. It supports both CLI and HTTP modes, making it useful in local testing, pipelines, and Kubernetes environments.

## Features

- 🔌 Run checks for:
  - DNS resolution
  - TCP connectivity
  - HTTP availability
- 🧪 Two modes of operation:
  - `cli` – Run one-off checks from the command line
  - `http` – Expose an HTTP server for programmatic checks and Kubernetes liveness/readiness probes
- 📦 Modular and extensible codebase
- ☁️ Built for future integration with Azure services:
  - Cosmos DB
  - SQL Database
  - Blob Storage
  - Managed Identity authentication

---

## Getting Started

### Prerequisites

- Go 1.21 or later

### Cloning the Repo

```bash
git clone https://github.com/BesQpin/orb.git
cd orb
go build ./cmd/orbmain
```

---

## Usage

### CLI Mode

Run the application in CLI mode with:

```bash
./orbmain --mode=cli
```

Then use subcommands:

#### DNS Check

```bash
./orbmain dns google.com
```

#### TCP Check

```bash
./orbmain tcp example.com:443
```

#### HTTP Check

```bash
./orbmain http https://example.com
```

---

### HTTP Server Mode

Run the HTTP server:

```bash
./orbmain --mode=http
```

Endpoints:

- `GET /healthz/live` – Liveness probe
- `GET /healthz/ready` – Readiness probe
- `GET /check/dns?host=google.com`
- `GET /check/tcp?addr=example.com:443`
- `GET /check/http?url=https://example.com`

---

## Project Structure

```
orb/
├── cmd/
│   └── orbmain.go        # Entry point
├── internal/
│   ├── cli/              # CLI setup and parsing
│   ├── server/           # HTTP server
│   └── checks/
│       ├── dns/          # DNS checks
│       ├── tcp/          # TCP checks
│       └── httpcheck/    # HTTP checks
```

---

## Extending Orb

Orb is built with modularity in mind. To add a new check:

1. Create a new package under `internal/checks/`
2. Implement both CLI and HTTP handler functions
3. Register them in `internal/cli/cli.go` and `internal/server/server.go`

---

## License

MIT © 2025 Peter Williams
