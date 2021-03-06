// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.13.0
// source: user.sql

package sqlc

import (
	"context"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users (addr, name, random_msg) VALUES ($1, $2, $3) RETURNING addr, admin, name, pfp, random_msg
`

type CreateUserParams struct {
	Addr      string `json:"addr"`
	Name      string `json:"name"`
	RandomMsg string `json:"randomMsg"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (Users, error) {
	row := q.db.QueryRowContext(ctx, createUser, arg.Addr, arg.Name, arg.RandomMsg)
	var i Users
	err := row.Scan(
		&i.Addr,
		&i.Admin,
		&i.Name,
		&i.Pfp,
		&i.RandomMsg,
	)
	return i, err
}

const getUser = `-- name: GetUser :one
SELECT addr, admin, name, pfp, random_msg FROM users WHERE addr = $1
`

func (q *Queries) GetUser(ctx context.Context, addr string) (Users, error) {
	row := q.db.QueryRowContext(ctx, getUser, addr)
	var i Users
	err := row.Scan(
		&i.Addr,
		&i.Admin,
		&i.Name,
		&i.Pfp,
		&i.RandomMsg,
	)
	return i, err
}

const updateUser = `-- name: UpdateUser :one
UPDATE users SET name = $2, pfp=$3, random_msg=$4 WHERE addr = $1 RETURNING addr, admin, name, pfp, random_msg
`

type UpdateUserParams struct {
	Addr      string      `json:"addr"`
	Name      string      `json:"name"`
	Pfp       interface{} `json:"pfp"`
	RandomMsg string      `json:"randomMsg"`
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) (Users, error) {
	row := q.db.QueryRowContext(ctx, updateUser,
		arg.Addr,
		arg.Name,
		arg.Pfp,
		arg.RandomMsg,
	)
	var i Users
	err := row.Scan(
		&i.Addr,
		&i.Admin,
		&i.Name,
		&i.Pfp,
		&i.RandomMsg,
	)
	return i, err
}
