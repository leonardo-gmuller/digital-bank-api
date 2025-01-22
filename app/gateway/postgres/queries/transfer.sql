-- name: ListTransfer :many
SELECT account.cpf as account_destination_cpf, amount, transfer.created_at FROM transfer
JOIN account ON account.id = account_destination_id
WHERE account_origin_id = $1
AND transfer.deleted_at IS NULL;

-- name: CreateTransfer :exec
INSERT INTO transfer (account_destination_id, account_origin_id, amount) 
VALUES ($1, $2, $3);