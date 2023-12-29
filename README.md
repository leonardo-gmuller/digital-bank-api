# GO Digital Bank API
This is a basic API that performs transactions between accounts.

## Built and Running with:
- Go
- Docker

## Folders Structure

📦src<br>
 ┣ 📂controllers<br>
 ┣ 📂database<br>
 ┣ 📂middleware<br>
 ┣ 📂models<br>
 ┣ 📂routes<br>
 ┗ 📜server.go<br>

 
`controllers`: methods for each endpoint

`database`: database related, including queries, connection, migrations

`middleware`: handlers for routes, for example: authentication

`models`: database entities, a mirror from the schema

`routes`: endpoints for API


## Endpoints
- `GET /accounts` - get a list of accounts
- `GET /accounts/{account_id}/balance` - get account balance
- `POST /accounts` - create an `Account`
- `POST /login` - authentic the user
- `GET /transfers` - gets the authenticated user's transfer list.
- `POST /transfers` - makes a transfer from one `Account` to another.

## Run Locally
First, set the environment variables by creating a .env in the root of the project (you can use `.env_template`).

Run `docker-compose up --build`

## Run Tests

Check if you have Postgres running on your machine and check the environment variables to connect to the database you want to run the tests on.

- For test account `go test -v src/tests/model_account_test.go`
- For test routes of accounts `go test -v src/tests/accounts_test.go`
- For test routes of transfers `go test -v src/tests/transfers_test.go`
- For test route of login `go test -v src/tests/login_test.go`