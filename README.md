# Personalidades API

> **Idioma:** [English](#personalidades-api) · [Português](README.pt-BR.md)

A simple CRUD API for managing "personalities" (name + history) built with **Go** and **PostgreSQL**, with a **React** frontend that lists the records.

## Features

- REST API with full CRUD (`GET`, `POST`, `PUT`, `DELETE`)
- PostgreSQL persistence via [GORM](https://gorm.io)
- JSON responses with consistent HTTP status codes and error handling
- CORS enabled for browser access
- React frontend consuming the API
- Configuration via environment variables
- `docker-compose` for a quick PostgreSQL setup

## Tech Stack

| Layer    | Technology                                   |
| -------- | -------------------------------------------- |
| Backend  | Go 1.19, [gorilla/mux](https://github.com/gorilla/mux), GORM |
| Database | PostgreSQL                                   |
| Frontend | React 17 (Create React App), axios           |

## Architecture

```
┌──────────────┐   HTTP/JSON    ┌──────────────┐   SQL    ┌──────────────┐
│   React SPA  │ ─────────────▶ │   Go API     │ ───────▶ │  PostgreSQL  │
│  (port 3000) │                │  (port 8000) │          │  (port 5432) │
└──────────────┘                └──────────────┘          └──────────────┘
```

- `routes` → HTTP routing and CORS
- `controllers` → request handlers (business logic)
- `models` → data models
- `database` → connection to PostgreSQL
- `middleware` → shared middleware

## Prerequisites

- [Go](https://go.dev/dl/) 1.19+
- [Node.js](https://nodejs.org/) 14+ (for the frontend)
- [PostgreSQL](https://www.postgresql.org/) (or Docker)

## Quick Start

### 1. Start PostgreSQL

Using Docker (recommended):

```bash
docker compose up -d
```

Or use an existing PostgreSQL instance and create the schema:

```bash
psql -U postgres -d postgres -f personalidade.sql
```

### 2. Run the API

```bash
cp .env.example .env   # optional: adjust DB credentials
go run main.go
```

The API starts at `http://localhost:8000`.

### 3. Run the frontend

```bash
cd frontend-personalidades
npm install
npm start
```

Open `http://localhost:3000`.

## Configuration

The API reads its configuration from environment variables. See [`.env.example`](.env.example) for the full list:

| Variable     | Default     | Description          |
| ------------ | ----------- | -------------------- |
| `DB_HOST`    | `localhost` | PostgreSQL host      |
| `DB_PORT`    | `5432`      | PostgreSQL port      |
| `DB_USER`    | `postgres`  | Database user        |
| `DB_PASSWORD`| `changeme`  | Database password    |
| `DB_NAME`    | `postgres`  | Database name        |
| `DB_SSLMODE` | `disable`   | SSL mode             |

The frontend reads `REACT_APP_API_URL` (defaults to `http://localhost:8000`) to locate the API.

## API Endpoints

Base URL: `http://localhost:8000`

| Method | Route                    | Description             |
| ------ | ------------------------ | ----------------------- |
| GET    | `/`                      | Home page               |
| GET    | `/api/health`            | Health check            |
| GET    | `/api/personalidades`    | List all personalities  |
| GET    | `/api/personalidades/{id}` | Get a single personality |
| POST   | `/api/personalidades`    | Create a personality    |
| PUT    | `/api/personalidades/{id}` | Update a personality  |
| DELETE | `/api/personalidades/{id}` | Delete a personality  |

The list endpoint supports pagination via `?page=` and `?limit=`. Create/update
validate that `nome` is not empty. See [`docs/api.md`](docs/api.md) for details.

> Detailed examples (requests, responses, status codes) are available in [`docs/api.md`](docs/api.md).

## Project Structure

```
go_api/
├── controllers/        # HTTP handlers
├── database/           # Database connection
├── middleware/         # Shared middleware
├── models/             # Data models
├── routes/             # Route definitions and CORS
├── frontend-personalidades/  # React frontend
├── docs/               # Detailed documentation
├── main.go             # Entry point
├── personalidade.sql   # Schema + seed data
├── docker-compose.yml  # PostgreSQL service
└── .env.example        # Environment variables template
```

## Documentation

- [Setup & installation](docs/setup.md)
- [API reference](docs/api.md)
- [Database](docs/database.md)

## Roadmap / Known Improvements

- Add automated tests (backend and frontend)
- Migrate the frontend to function components / hooks (or a newer toolchain)
- Restrict CORS to specific origins instead of `*`
- Add request validation and pagination

## License

This project is licensed under the [MIT License](LICENSE).
