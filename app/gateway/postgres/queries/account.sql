-- name: ListAccounts :many
SELECT * 
FROM account
WHERE deleted_at IS NULL;

-- name: GetAccountById :one
SELECT * 
FROM account
WHERE id = $1
AND deleted_at IS NULL;

-- name: GetAccountByCpf :one
SELECT * 
FROM account
WHERE cpf = $1
AND deleted_at IS NULL;

-- name: GetAccountBalanceById :one
SELECT name, balance 
FROM account
WHERE id = $1
AND deleted_at IS NULL;

-- name: UpdateBalance :exec
UPDATE account SET balance = $1
WHERE id = $2
AND deleted_at IS NULL;

-- name: CreateAccount :exec
INSERT INTO account (name, cpf, secret, balance) 
VALUES ($1, $2, $3, $4);