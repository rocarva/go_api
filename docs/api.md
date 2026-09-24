# API Reference

Base URL: `http://localhost:8000`

All endpoints return `application/json`. Error responses follow the shape:

```json
{ "erro": "<message>" }
```

## Resource: Personalidade

```json
{
  "id": 1,
  "nome": "Deodato Petit Wertheimer",
  "historia": "Deodato Petit Wertheimer foi um médico e político brasileiro..."
}
```

| Field      | Type    | Description        |
| ---------- | ------- | ------------------ |
| `id`       | integer | Auto-generated ID  |
| `nome`     | string  | Name               |
| `historia` | string  | History/description|

## Endpoints

### `GET /`

Returns a plain text home page.

```
HTTP/1.1 200 OK
Home Page
```

### `GET /api/health`

Health check endpoint.

**Response** — `200 OK`:

```json
{ "status": "ok" }
```

### `GET /api/personalidades`

Lists all personalities.

**Query parameters** (both optional; pagination is enabled when either is present):

| Param   | Type    | Default | Description              |
| ------- | ------- | ------- | ------------------------ |
| `page`  | integer | `1`     | Page number (1-based)    |
| `limit` | integer | `10`    | Records per page         |

When paginating, the response includes the `X-Total-Count` header with the
total number of records.

**Response** — `200 OK`:

```json
[
  { "id": 1, "nome": "...", "historia": "..." },
  { "id": 2, "nome": "...", "historia": "..." }
]
```

**Example (paginated):**

```bash
curl "http://localhost:8000/api/personalidades?page=1&limit=10"
```

### `GET /api/personalidades/{id}`

Returns a single personality.

**Response** — `200 OK`:

```json
{ "id": 1, "nome": "...", "historia": "..." }
```

**Errors:**

| Status | When            |
| ------ | --------------- |
| `404`  | ID not found    |
| `500`  | Database error  |

### `POST /api/personalidades`

Creates a new personality.

**Request body** (JSON):

```json
{ "nome": "Novo Nome", "historia": "Nova história" }
```

The `nome` field is required (non-empty).

**Response** — `201 Created` (the created record, including its `id`).

**Errors:**

| Status | When                             |
| ------ | -------------------------------- |
| `400`  | Invalid JSON body or empty `nome`|
| `500`  | Database error                   |

**Example:**

```bash
curl -X POST http://localhost:8000/api/personalidades \
  -H "Content-Type: application/json" \
  -d '{"nome":"Exemplo","historia":"História de exemplo"}'
```

### `PUT /api/personalidades/{id}`

Updates an existing personality.

**Request body** (JSON) — fields to update:

```json
{ "nome": "Nome Atualizado", "historia": "História atualizada" }
```

The `nome` field, when provided, cannot be empty.

**Response** — `200 OK` (the updated record).

**Errors:**

| Status | When                             |
| ------ | -------------------------------- |
| `400`  | Invalid JSON body or empty `nome`|
| `404`  | ID not found                     |
| `500`  | Database error                   |

**Example:**

```bash
curl -X PUT http://localhost:8000/api/personalidades/1 \
  -H "Content-Type: application/json" \
  -d '{"nome":"Nome Atualizado","historia":"História atualizada"}'
```

### `DELETE /api/personalidades/{id}`

Deletes a personality.

**Response** — `200 OK` (the deleted record).

**Errors:**

| Status | When            |
| ------ | --------------- |
| `404`  | ID not found    |
| `500`  | Database error  |

**Example:**

```bash
curl -X DELETE http://localhost:8000/api/personalidades/1
```
