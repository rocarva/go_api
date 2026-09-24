# Personalidades API

> **Idioma:** [English](README.md) · [Português](#personalidades-api)

Uma API CRUD simples para gerenciar "personalidades" (nome + história) construída com **Go** e **PostgreSQL**, com um frontend em **React** que lista os registros.

## Funcionalidades

- API REST com CRUD completo (`GET`, `POST`, `PUT`, `DELETE`)
- Persistência em PostgreSQL via [GORM](https://gorm.io)
- Respostas JSON com códigos de status HTTP consistentes e tratamento de erros
- CORS habilitado para acesso via navegador
- Frontend React consumindo a API
- Configuração via variáveis de ambiente
- `docker-compose` para subir o PostgreSQL rapidamente

## Stack

| Camada   | Tecnologia                                    |
| -------- | --------------------------------------------- |
| Backend  | Go 1.19, [gorilla/mux](https://github.com/gorilla/mux), GORM |
| Banco    | PostgreSQL                                    |
| Frontend | React 17 (Create React App), axios            |

## Arquitetura

```
┌──────────────┐   HTTP/JSON    ┌──────────────┐   SQL    ┌──────────────┐
│   React SPA  │ ─────────────▶ │   Go API     │ ───────▶ │  PostgreSQL  │
│  (porta 3000)│                │  (porta 8000)│          │  (porta 5432)│
└──────────────┘                └──────────────┘          └──────────────┘
```

- `routes` → roteamento HTTP e CORS
- `controllers` → handlers das requisições (regras de negócio)
- `models` → modelos de dados
- `database` → conexão com o PostgreSQL
- `middleware` → middleware compartilhado

## Pré-requisitos

- [Go](https://go.dev/dl/) 1.19+
- [Node.js](https://nodejs.org/) 14+ (para o frontend)
- [PostgreSQL](https://www.postgresql.org/) (ou Docker)

## Início rápido

### 1. Suba o PostgreSQL

Usando Docker (recomendado):

```bash
docker compose up -d
```

Ou use uma instância existente do PostgreSQL e crie o schema:

```bash
psql -U postgres -d postgres -f personalidade.sql
```

### 2. Rode a API

```bash
cp .env.example .env   # opcional: ajuste as credenciais do banco
go run main.go
```

A API inicia em `http://localhost:8000`.

### 3. Rode o frontend

```bash
cd frontend-personalidades
npm install
npm start
```

Abra `http://localhost:3000`.

## Configuração

A API lê sua configuração por variáveis de ambiente. Veja [`.env.example`](.env.example) para a lista completa:

| Variável      | Padrão      | Descrição           |
| ------------- | ----------- | ------------------- |
| `DB_HOST`     | `localhost` | Host do PostgreSQL  |
| `DB_PORT`     | `5432`      | Porta do PostgreSQL |
| `DB_USER`     | `postgres`  | Usuário do banco    |
| `DB_PASSWORD` | `changeme`  | Senha do banco      |
| `DB_NAME`     | `postgres`  | Nome do banco       |
| `DB_SSLMODE`  | `disable`   | Modo SSL            |

O frontend lê `REACT_APP_API_URL` (padrão `http://localhost:8000`) para localizar a API.

## Endpoints da API

URL base: `http://localhost:8000`

| Método | Rota                      | Descrição                |
| ------ | ------------------------- | ------------------------ |
| GET    | `/`                       | Página inicial           |
| GET    | `/api/personalidades`     | Lista todas as personalidades |
| GET    | `/api/personalidades/{id}` | Retorna uma personalidade |
| POST   | `/api/personalidades`     | Cria uma personalidade   |
| PUT    | `/api/personalidades/{id}` | Atualiza uma personalidade |
| DELETE | `/api/personalidades/{id}` | Deleta uma personalidade |

> Exemplos detalhados (requisições, respostas, códigos de status) em [`docs/api.md`](docs/api.md).

## Estrutura do projeto

```
go_api/
├── controllers/        # Handlers HTTP
├── database/           # Conexão com o banco
├── middleware/         # Middleware compartilhado
├── models/             # Modelos de dados
├── routes/             # Definição de rotas e CORS
├── frontend-personalidades/  # Frontend React
├── docs/               # Documentação detalhada
├── main.go             # Ponto de entrada
├── personalidade.sql   # Schema + dados de exemplo
├── docker-compose.yml  # Serviço do PostgreSQL
└── .env.example        # Modelo de variáveis de ambiente
```

## Documentação

- [Instalação e configuração](docs/setup.md)
- [Referência da API](docs/api.md)
- [Banco de dados](docs/database.md)

## Melhorias futuras / conhecidas

- Adicionar testes automatizados (backend e frontend)
- Migrar o frontend para componentes de função / hooks (ou toolchain mais novo)
- Restringir o CORS a origens específicas em vez de `*`
- Adicionar validação de entrada e paginação

## Licença

Este projeto está licenciado sob a [Licença MIT](LICENSE).
