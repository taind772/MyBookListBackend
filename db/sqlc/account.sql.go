// Code generated by sqlc. DO NOT EDIT.
// source: account.sql

package db

import (
	"context"
)

const createAccount = `-- name: CreateAccount :one
INSERT INTO accounts (
    id,
    username,
    email,
    role
) VALUES (
    $1, $2, $3, $4
) RETURNING id, username, email, avatar_uri, role, modified_at, created_at
`

type CreateAccountParams struct {
	ID       int64
	Username string
	Email    string
	Role     AccountRoles
}

func (q *Queries) CreateAccount(ctx context.Context, arg CreateAccountParams) (Account, error) {
	row := q.db.QueryRowContext(ctx, createAccount,
		arg.ID,
		arg.Username,
		arg.Email,
		arg.Role,
	)
	var i Account
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Email,
		&i.AvatarUri,
		&i.Role,
		&i.ModifiedAt,
		&i.CreatedAt,
	)
	return i, err
}

const getAccount = `-- name: GetAccount :one
SELECT id, username, email, avatar_uri, role, modified_at, created_at
FROM accounts
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetAccount(ctx context.Context, id int64) (Account, error) {
	row := q.db.QueryRowContext(ctx, getAccount, id)
	var i Account
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Email,
		&i.AvatarUri,
		&i.Role,
		&i.ModifiedAt,
		&i.CreatedAt,
	)
	return i, err
}

const listAccounts = `-- name: ListAccounts :many
SELECT id, username, email, avatar_uri, role, modified_at, created_at
FROM accounts
LIMIT $1
OFFSET $2
`

type ListAccountsParams struct {
	Limit  int32
	Offset int32
}

func (q *Queries) ListAccounts(ctx context.Context, arg ListAccountsParams) ([]Account, error) {
	rows, err := q.db.QueryContext(ctx, listAccounts, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Account
	for rows.Next() {
		var i Account
		if err := rows.Scan(
			&i.ID,
			&i.Username,
			&i.Email,
			&i.AvatarUri,
			&i.Role,
			&i.ModifiedAt,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}