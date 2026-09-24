# Setup & Installation

This guide covers how to run the API and the frontend, either locally or with Docker.

## Prerequisites

| Tool          | Minimum version | Purpose                  |
| ------------- | --------------- | ------------------------ |
| Go            | 1.19            | Build/run the backend    |
| Node.js       | 14              | Run the React frontend   |
| Docker        | —               | (Optional) PostgreSQL    |
| PostgreSQL    | 12+             | Database (if no Docker)  |

## 1. Database

### Option A — Docker (recommended)

From the project root:

```bash
docker compose up -d
```

This starts a PostgreSQL 15 container on port `5432` with the credentials
defined in [`docker-compose.yml`](../docker-compose.yml) (user `postgres`,
password `changeme`, database `postgres`).

To stop it:

```bash
docker compose down
```

To stop and remove the data volume:

```bash
docker compose down -v
```

### Option B — Existing PostgreSQL

Create the schema and seed data:

```bash
psql -U postgres -d postgres -f personalidade.sql
```

## 2. Backend (Go API)

1. Create your environment file (optional — sensible defaults are used if omitted):

   ```bash
   cp .env.example .env
   ```

   > The application reads variables from the environment (not from the file).
   > On Linux/macOS you can source the file: `export $(cat .env | xargs)`.
   > On Windows PowerShell, set each variable with `$env:DB_HOST = "localhost"`, etc.

2. Run the server:

   ```bash
   go run main.go
   ```

3. Confirm it is up:

   ```bash
   curl http://localhost:8000/api/personalidades
   ```

## 3. Frontend (React)

```bash
cd frontend-personalidades
npm install
npm start
```

The app runs at `http://localhost:3000`.

### Pointing the frontend at a different API

By default the frontend calls `http://localhost:8000`. To override it, set the
`REACT_APP_API_URL` environment variable before starting:

```bash
# Linux/macOS
REACT_APP_API_URL=http://api.example.com npm start

# Windows PowerShell
$env:REACT_APP_API_URL = "http://api.example.com"; npm start
```

## Environment variables

### Backend

| Variable     | Default     | Description          |
| ------------ | ----------- | -------------------- |
| `DB_HOST`    | `localhost` | PostgreSQL host      |
| `DB_PORT`    | `5432`      | PostgreSQL port      |
| `DB_USER`    | `postgres`  | Database user        |
| `DB_PASSWORD`| `changeme`  | Database password    |
| `DB_NAME`    | `postgres`  | Database name        |
| `DB_SSLMODE` | `disable`   | SSL mode             |

### Frontend

| Variable            | Default                 | Description        |
| ------------------- | ----------------------- | ------------------ |
| `REACT_APP_API_URL` | `http://localhost:8000` | API base URL       |

## Troubleshooting

- **`log.Panic: Erro ao conectar ao banco de dados`** — check that PostgreSQL
  is running and that `DB_HOST`/`DB_PORT`/`DB_USER`/`DB_PASSWORD` match your
  instance.
- **Frontend shows no data** — verify the API is up and that `REACT_APP_API_URL`
  points to it. Also check CORS (the API currently allows all origins).
- **Port already in use** — the API listens on `:8000`. Change it in
  `routes/routes.go` if needed.
