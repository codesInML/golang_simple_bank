-- name: CreateAccount :one
INSERT INTO accounts (
  owner, balance, currency
) VALUES (
  ?, ?, ?
) RETURNING *;

-- name: GetAccount :one
SELECT * FROM accounts
WHERE id = ? LIMIT 1;

-- name: ListAccounts :many
SELECT * FROM accounts
ORDER BY id
LIMIT ?
OFFSET ?;

-- name: UpdateAccount :exec
UPDATE accounts SET balance = $2
WHERE id = $1;