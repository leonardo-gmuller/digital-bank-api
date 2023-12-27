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


### Run locally
`docker-compose build`
`docker-compose up`