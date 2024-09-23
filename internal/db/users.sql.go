// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: users.sql

package db

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users(
    id,
    created_at,
    updated_at,
    name,
    phoneNumber,
    email,
    address,
    password,
    ApiKey
)
VALUES ($1,$2,$3,$4,$5,$6,$7,$8,encode(sha256(random()::text::bytea),'hex'))
RETURNING id, created_at, updated_at, name, phonenumber, email, address, password, apikey
`

type CreateUserParams struct {
	ID          uuid.UUID
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Name        string
	Phonenumber string
	Email       string
	Address     string
	Password    string
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, createUser,
		arg.ID,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.Name,
		arg.Phonenumber,
		arg.Email,
		arg.Address,
		arg.Password,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Name,
		&i.Phonenumber,
		&i.Email,
		&i.Address,
		&i.Password,
		&i.Apikey,
	)
	return i, err
}

const getUser = `-- name: GetUser :one
SELECT id, created_at, updated_at, name, phonenumber, email, address, password, apikey FROM users WHERE ApiKey = $1
`

func (q *Queries) GetUser(ctx context.Context, apikey string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUser, apikey)
	var i User
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Name,
		&i.Phonenumber,
		&i.Email,
		&i.Address,
		&i.Password,
		&i.Apikey,
	)
	return i, err
}
