# Copilot Instructions for fantasy-products

## Project Overview
- This is a Go backend for a fantasy store, organized in a clean architecture style with clear separation between domain, repository, service, handler, and application layers.
- Data is persisted in MySQL, with schema/data initialization via Docker Compose and SQL files in `docs/db/mysql/`.
- The project supports both RESTful API endpoints (using chi router) and bulk import endpoints for JSON data.

## Key Directories & Files
- `cmd/main.go`: Application entrypoint, configures logging, loads env vars, and starts the server.
- `internal/application/`: App setup, dependency injection, and route registration.
- `internal/domain/`: Core business entities and interfaces (e.g., `Customer`, `Product`).
- `internal/repository/`: MySQL implementations for each entity.
- `internal/service/`: Business logic, often wraps repository.
- `internal/handler/`: HTTP handlers, one per entity, with REST and bulk endpoints.
- `utils/readJson.go`: Generic loader for JSON data, used for bulk import endpoints.
- `docs/db/json/`: Example data for bulk import (e.g., `customers.json`).
- `docker-compose.yml`: Defines app and db containers, loads env vars, mounts SQL/JSON files.

## Patterns & Conventions
- **Dependency Injection:** All handlers/services/repositories are wired in `application_default.go`.
- **REST Endpoints:** Standard CRUD at `/entity` (GET, POST), plus `/entity/bulk` for JSON import, and `/entity/stats/*` for analytics.
- **Bulk Import:** Use `POST /entity/bulk` with a JSON array or load from file using `utils.ReadJson`.
- **Logging:** Use `log/slog` for structured logs. Avoid `fmt.Println` in production code.
- **Environment:** All config (DB, API key, etc.) via `.env` and referenced in `docker-compose.yml`.
- **Testing:** Uses [txdb](https://github.com/DATA-DOG/go-txdb) for transactional DB tests. See `tests/txdb.go` for setup.

## Developer Workflows
- **Build/Run:** Use `docker-compose up --build` to start app and DB with correct env/config.
- **Test:** Run Go tests with `go test ./...` (DB tests use txdb and `.env` config).
- **Debug:** Logs are output in JSON via slog; check container logs for issues.
- **DB Reset:** To re-run SQL init scripts, use `docker-compose down -v && docker-compose up`.

## Integration Points
- **MySQL:** All persistence via MySQL, DSN built with `mysql.Config.FormatDSN()`.
- **txdb:** For tests, DB is registered with txdb using the same DSN logic as production.
- **Env Loading:** Use `github.com/joho/godotenv` for loading `.env` in tests and local runs.

## Examples
- See `internal/handler/customer.go` for REST and bulk import handler patterns.
- See `utils/readJson.go` for generic JSON loader usage.
- See `tests/txdb.go` for DB test setup and DSN construction.

---
If any section is unclear or missing key project knowledge, please provide feedback for improvement.
