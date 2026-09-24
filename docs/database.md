# Database

The API uses **PostgreSQL** and accesses it through [GORM](https://gorm.io).

## Schema

The schema is defined in [`personalidade.sql`](../personalidade.sql) at the
project root:

```sql
create table personalidades(
    id serial primary key,
    nome varchar,
    historia varchar
);
```

| Column    | Type         | Notes                     |
| --------- | ------------ | ------------------------- |
| `id`      | `serial`     | Primary key, auto-increment |
| `nome`     | `varchar`   | Personality name          |
| `historia` | `varchar`   | Personality history       |

## Model (Go)

The Go model lives in [`models/models.go`](../models/models.go):

```go
type Personalidade struct {
	Id       int    `json:"id"`
	Nome     string `json:"nome"`
	Historia string `json:"historia"`
}
```

## Connection

The connection is built in [`database/db.go`](../database/db.go) from
environment variables (with defaults). Example connection string:

```
host=localhost user=postgres password=changeme dbname=postgres port=5432 sslmode=disable
```

## Applying the schema

```bash
psql -U postgres -d postgres -f personalidade.sql
```

> Note: `personalidade.sql` contains `INSERT` statements with two sample
> records used for demo purposes.

## Seed data

The seed data includes two Brazilian historical figures:

1. **Deodato Petit Wertheimer** — physician and politician.
2. **Carmela Dutra** — First Lady of Brazil (1946–1947).

## Migrations

This project does **not** use a migration tool (e.g., GORM AutoMigrate or
golang-migrate). The schema is applied manually via the SQL script. If you need
versioned migrations, consider introducing `gormigrate` or `golang-migrate`.
