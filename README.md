# GO Digital Bank API
This is a basic API that performs transactions between accounts.

## Built and running with:
- Go
- Docker

## Folders structure

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


## Endpoints description
- `GET /accounts` - get a list of accounts
- `GET /accounts/{account_id}/balance` - get account balance
- `POST /accounts` - create an `Account`
- `POST /login` - authentic the user
- `GET /transfers` - gets the authenticated user's transfer list.
- `POST /transfers` - makes a transfer from one `Account` to another.

## Run locally
`docker-compose build`
`docker-compose up`