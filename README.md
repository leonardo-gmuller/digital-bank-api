# Digital-Bank-API

Digital-Bank-API is a example of web application designed to manage banking operations, clients, transactions, and ensure secure access. The application follows a *clean architecture* approach to ensure maintainability and scalability.

## Tech Stack

- **Go**: Backend is built using Go.
- **Chi Router**: Lightweight and fast HTTP router.
- **Golang-Migrate**: For database migrations.
- **JWT**: Authentication middleware using JSON Web Tokens.



## Folders Structure

- `build`: files for build
    - `docker-compose.yaml`
    - `Dockerfile`
- `app`
    - `cmd`
        - `server`: main file
    - `config`: configure environment
    - `domain`: domain related, including entities, usecases, dto
    - `gateway`
        - `api`: api handlers, middlewares
        - `postgres`: DB connection and repositories
    - `pkg`: open library
    - `resources`: resources for API
    - `tests`
        - `mocks`: mocks for tests
- `local`
    - `docker-compose.yaml`: For development
- `tests`: http files to help run tests


 

## Installation
1. Clone o repositório:
   ```bash
   git clone github.com/leonardo-gmuller/digital-bank-api
   cd digital-bank-api
   ```

2. Install dependencies
    ```bash
    go mod tidy
    ```

## Start
#### Requirements
Make sure you have the following software installed:

- [Go](https://go.dev/doc/install)<br>
- [Docker](https://www.docker.com/)
- [Docker Compose](https://docs.docker.com/compose/)
- [Golang-migrate](https://github.com/golang-migrate/migrate/)
- [Makefile](https://www.gnu.org/software/make/manual/make.html)

To help with the startup, I recommend you use make:
- For development:
```bash
make start-local
```

- For build:
```bash
make start-build
```

## Endpoints
Prefix: `http://localhost:8000`
- `GET /accounts` - get a list of accounts
- `GET /accounts/{account_id}/balance` - get account balance
- `POST /accounts` - create an `Account`
- `POST /auth` - authentic the user
- `GET /transfers` - gets the authenticated user's transfer list.
- `POST /transfers` - makes a transfer from one `Account` to another.


