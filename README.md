# To-Do gRPC Service

A gRPC service for task management and user authentication built with Go. Designed around clean architecture principles, jwt-tokens security, and schema-first request validation.

---

## Tech Stack

* **Language:** Go 1.27
* **API Framework:** gRPC, Protocol Buffers (proto3)
* **Database:** PostgreSQL
* **DB Tooling:** `sqlc` (type-safe SQL generation)
* **Validation:** `buf.build/go/protovalidate` (Interceptor-based request validation)
* **Security:** Argon2id (`golang.org/x/crypto/argon2`), JWT (`golang-jwt/jwt/v5`)
* **Infrastructure:** Docker, Docker Compose, Adminer

---

## Key Features & Architecture

* **Schema-First API Design:** Complete separation of API contracts (`.proto` files) from implementation logic.
* **gRPC Interceptor Pipeline:** 
  * **Auth Interceptor:** Validates Bearer tokens from incoming `metadata` headers on protected routes.
  * **Validation Interceptor:** Automatic request payload validation via `protovalidate` before reaching service handlers.
* **Robust Authentication:** Password hashing using the Argon2id algorithm with unique salting and constant-time hash comparison.
* **Clean Persistence Layer:** Type-safe database queries generated via `sqlc` without heavy ORM overhead.

---

## Getting Started

### Prerequisites

* Docker and Docker Compose
* Go 1.22+ (for local development)
* `protoc` compiler with Go plugins (optional, for code generation)

### 1. Environment Configuration

Copy and fill out the example environment file:

```bash
cp .env.example .env
```

### 2. Run with Docker Compose
Spin up the entire stack including the Go application, PostgreSQL database, and Adminer web UI. Database migrations located in ./migrations are automatically executed on startup:

```bash
docker-compose up -d --build
```

### 3. Protobuf Code Generation
If you modify the .proto schemas, run the generation script:

```bash
chmod +x generate.sh
./generate.sh
```

---

## API Specification

### UserService
Handles user identity and authentication. All endpoints are public (no authentication token required).

* **Register** - Registers a new user account. Hashes passwords using Argon2id and stores user details.
* **Login** - Authenticates user credentials and returns a signed JWT access token.

### TaskService
Provides task management capabilities. Every endpoint requires a valid JWT sent in gRPC Metadata:
authorization: Bearer <token>

* **CreateTask** - Creates a new task bound to the authenticated user.
* **GetTaskByID** - Fetches details of a specific task by its ID.
* **UpdateTask** - Updates fields on an existing task.
* **DeleteTask** - Removes a task by ID.
* **GetTasksByFilter** - Retrieves a paginated, filtered, and sorted list of tasks owned by the authenticated user.

---