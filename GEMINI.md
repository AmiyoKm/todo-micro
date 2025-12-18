# Project Context: Todo Microservices

## Overview
This is a full-stack microservices application named `todo-micro`, designed to manage todo lists with user authentication. The project uses a **Go** backend with a **React** frontend, orchestrated via **Docker Compose**.

## Architecture
The system consists of the following services:

### Frontend
- **Path:** `frontend/`
- **Tech Stack:** React, Vite, TypeScript, Tailwind CSS, Shadcn UI, React Query.
- **Port:** `5173` (mapped to host).
- **Description:** The user interface for the application.

### Backend Services
1.  **API Gateway**
    - **Path:** `api-gateway/`
    - **Tech Stack:** Go (Standard Library `net/http`).
    - **Port:** `3000` (mapped to host).
    - **Description:** The entry point for all client requests. Handles routing, authentication middleware, and communicates with internal services via gRPC.
    - **Key Routes:**
        - `POST /register`, `POST /login` (Public)
        - `GET/POST/PATCH/DELETE /todos`, `GET /users/me` (Protected)

2.  **User Service**
    - **Path:** `user/`
    - **Tech Stack:** Go, gRPC, Postgres (`user-db`).
    - **Description:** Manages user registration, login, and profile data.
    - **Database:** `user_micro` (Postgres 18).

3.  **Todo Service**
    - **Path:** `server/`
    - **Tech Stack:** Go, gRPC, Postgres (`todo-db`), Redis.
    - **Description:** Manages todo items (CRUD).
    - **Database:** `todo_micro` (Postgres 18).
    - **Cache:** Redis (for caching todos).

4.  **Mutate Todo**
    - **Path:** `mutate-todo/`
    - **Tech Stack:** Go, RabbitMQ.
    - **Description:** A worker service that likely processes asynchronous tasks or updates related to todos using RabbitMQ.

### Infrastructure
- **Databases:** PostgreSQL 18 (Alpine). Separate instances for `user` and `todo` services.
- **Cache:** Redis (Alpine).
- **Message Broker:** RabbitMQ (Management Alpine). Ports `5672` (AMQP) and `15672` (Management UI).

## Development & Usage

### Running the Project
The entire stack is containerized. To start all services:
```bash
docker compose up --build
```
*   **Frontend:** http://localhost:5173
*   **API Gateway:** http://localhost:3000
*   **RabbitMQ Management:** http://localhost:15672 (User: `guest`, Pass: `guest`)

### Service-Specific Commands

#### Frontend (`frontend/`)
- **Install Dependencies:** `npm install`
- **Run Dev Server:** `npm run dev`
- **Build:** `npm run build`
- **Lint:** `npm run lint`

#### Backend (`server/` and `user/`)
Both services use a `Makefile` for common tasks:
- **Run Locally:** `make run/api`
- **Apply Migrations:** `make db/migrations/up`
- **Create Migration:** `make db/migrations/new name=<name>`
- **Rollback Migration:** `make db/migrations/down`

### Key Configuration Files
- `docker-compose.yaml`: Service orchestration and environment variables.
- `proto/*.proto`: gRPC service definitions (`task.proto`, `user.proto`).
- `api-gateway/main.go`: Gateway routing and middleware logic.
- `server/internal/migrations/`: SQL migration files for the Todo service.
- `user/internal/migrations/`: SQL migration files for the User service.

## Conventions
- **API Style:** RESTful JSON API exposed by the Gateway, mapped to internal gRPC calls.
- **Database:** `golang-migrate` is used for schema versioning.
- **Environment:** Configuration is handled via environment variables (seen in `docker-compose.yaml` and `config.go` files).
