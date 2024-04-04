# GO Digital Bank API
This is a basic API that performs transactions between accounts.

## Built and Running with:
- Go
- Docker

## Folders Structure
📦build<br>
┣ 🐳docker-compose.yml<br>
┗ 🐳Dockerfile<br>
📦src<br>
 ┣ 📂app<br>
   ┣ 📂config<br>
   ┣ 📂domain<br>
   ┣ 📂gateway<br>
   ┣ 📂helpers<br>
   ┣ 📂resource<br>
 ┗ 📜server.go<br>

 
`config`: configure environment

`domain`: domain related, including entities, usecases, dto

`gateway`: api handlers, middlewares and repositories

`helpers`: functions helpers

`resource`: resources for API


## Endpoints
- `GET /accounts` - get a list of accounts
- `GET /accounts/{account_id}/balance` - get account balance
- `POST /accounts` - create an `Account`
- `POST /auth` - authentic the user
- `GET /transfers` - gets the authenticated user's transfer list.
- `POST /transfers` - makes a transfer from one `Account` to another.

## Run Locally
First, set the environment variables (you can use `.env_template`), don't forget to configure a Postgres database.


## Build
 Go to build folder:

 `cd build`

 Run docker-compose with build flag:

 `docker-compose up -d --build`


