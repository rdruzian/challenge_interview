# challenge_interview

A Go (Gin) REST API for a device CRUD challenge, using GORM and PostgreSQL. It includes Swagger documentation, automatic DB migrations at startup, and a Docker Compose setup for local development.

## Overview
- Language: Go 1.25 (multi-stage Docker build)
- Web framework: Gin
- ORM: GORM (PostgreSQL driver)
- Docs: Swagger (served via `gin-swagger`)
- DB: PostgreSQL 16 (alpine)


## Project Structure
```
.
├── Dockerfile
├── docker-compose.yml
├── docs/                # Generated Swagger docs (docs.go, swagger.json, swagger.yaml)
├── database/            # DB bootstrap + migrations
├── inbound/             # Service inbound interface
├── outbound/            # Repository outbound interface
├── middleware/          # Error and not found handlers
├── model/               # Domain models (e.g., Device)
├── repository/          # GORM repository (CRUD, to be implemented)
├── router/              # HTTP handlers + Swagger annotations
├── server/              # Server wiring (DB, repo, service, routes)
└── main.go              # Entry point and SwaggerInfo metadata
```

## Prerequisites
- Docker Desktop running (Windows/macOS/Linux)
- Docker Compose v2

## Quick Start (Docker)
1. Build and start services:
   - `docker compose up -d`
2. Verify containers:
   - `docker compose ps`
   - `docker compose logs -f app`
3. Access the API docs:
   - Swagger UI: `http://localhost:8080/swagger/index.html`
4. Stop services:
   - `docker compose down`
5. Reset the database (remove volumes):
   - `docker compose down -v`

## Services & Ports
- App (`challenge_app`):
  - Listens on `8080` (mapped to host `8080`)
- DB (`challenge_db`):
  - Container port `5432` mapped to host `25432`

## Configuration (Environment Variables)
The application reads DB settings from env vars (with defaults), and the compose file sets them appropriately:
- `DB_HOST` (default: `localhost` in Dockerfile, `db` in compose)
- `DB_PORT` (default: `25432` in Dockerfile, `5432` in compose)
- `DB_USER` (default: `admin`)
- `DB_NAME` (default: `challenge_db`)
- `DB_PASSWORD` (default: `123456`)

These defaults are compatible with `docker-compose.yml`. Adjust values in the compose file if needed.

## Health Checks
- App health (compose): access to Swagger UI `GET /swagger/index.html`.
- DB health (compose): `pg_isready -U admin -d challenge_db`.

## Swagger API Documentation
- Swagger UI is served at `GET /swagger/*any` (e.g., `/swagger/index.html`).
- Documented routes include:
  - `GET /api/v1/device` — list all devices
  - `GET /api/v1/device/{id}` — fetch a device by ID
  - `GET /api/v1/device/brand/{brand}` — list devices by brand
  - `GET /api/v1/device/state/{state}` — list devices by state
  - `POST /api/v1/device/create` — create a device
  - `DELETE /api/v1/device/{id}` — delete a device by ID

Note: The repository and service layers are implemented with GORM and state validation (allowed: `available`, `in-use`, `inactive`). Swagger UI and all CRUD endpoints are functional when the database is up.

### Regenerating Swagger Docs (optional, for local dev)
- Install the generator: `go install github.com/swaggo/swag/cmd/swag@v1.8.12`
- Generate docs: `swag init -g main.go -o docs`
- The app imports `github.com/rdruzian/challenge_interview/docs` and sets `SwaggerInfo` at runtime.

## Local Development (without Docker)
If you want to preview Swagger without DB, you can run in Swagger-only mode:
- PowerShell: `Set-Item -Path Env:SWAG_ONLY -Value "true"; go run .`
- Visit: `http://localhost:8080/swagger/index.html`

Normal mode (requires a reachable PostgreSQL):
- `go build ./...`
- `go run .`

## Troubleshooting
- Docker not running / image pull errors: ensure Docker Desktop is started.
- App cannot connect to DB:
  - Wait for `db` healthcheck to pass; the app depends on `db`.
  - Confirm host port `25432` is free and mapped correctly.
- Port `8080` in use: change the host port mapping in `docker-compose.yml` or adjust server settings.
- Swagger not loading: confirm the app is up and check logs.
- Windows GOPATH/GOROOT warning in logs can be ignored for this project setup.

## Useful Commands
- Build images: `docker compose build`
- Start services: `docker compose up -d`
- Follow logs: `docker compose logs -f app`
- Stop services: `docker compose down`
- Reset DB volume: `docker compose down -v`
- Check container health/status: `docker compose ps`

## Notes
- DB migrations run automatically at app startup (`database/migrations`).
- Container names: `challenge_db` (PostgreSQL), `challenge_app` (Go API).
- Runtime image: Alpine 3.19 with `curl` and `ca-certificates` for healthchecks.