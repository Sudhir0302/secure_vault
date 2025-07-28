# Secure Vault

Secure Vault is a secure file storage and sharing service built with a microservices architecture. It includes user authentication, file storage, and secure file sharing functionality.

## Tech Stack

* **Backend Language**: [Go (Golang)](https://golang.org/)
* **Web Framework**: [Gin Gonic](https://gin-gonic.com/)
* **ORM**: [GORM](https://gorm.io/)
* **Database**: [PostgreSQL](https://www.postgresql.org/)

## Services

This project is divided into three main microservices:

### 1. Auth Service (`localhost:8080`)

Handles user authentication and registration.

* `GET /test` – Health check
* `POST /signup` – Register a new user
* `POST /signin` – Login with email and password

### 2. Storage Service (`localhost:8081`)

Handles file uploads and retrievals.

* `GET /test` – Health check
* `POST /upload` – Upload a file (FormData: `file`, `userid`, `filename`)
* `GET /api/getfile` – Retrieve file (requires Bearer token)

### 3. Share Service (`localhost:8082`)

Manages secure file sharing with expiry and download limits.

* `GET /test` – Health check
* `POST /api/addshare` – Share a file with a password and limits
* `GET /api/getshare` – Access a shared file (requires Bearer token)

## Postman Collection

You can test all API endpoints using the provided [Postman collection](https://github.com/Sudhir0302/secure_vault).

## Getting Started

1. **Clone the repository:**

```bash
git clone https://github.com/Sudhir0302/secure_vault.git
cd secure_vault
```

2. **Set up PostgreSQL** and update DB credentials in each service.

3. **Run services:**

Each service can be run independently:

```bash
go run main.go
```

(Repeat this inside each service directory.)

## License

This project is licensed under the **MIT License**. See the [LICENSE](#license) section below for details.
