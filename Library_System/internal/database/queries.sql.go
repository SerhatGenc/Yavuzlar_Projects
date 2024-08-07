// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: queries.sql

package database

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const addBook = `-- name: AddBook :exec
INSERT INTO books (bookname, writer, user_id, deadline)
VALUES ($1, $2, $3, $4)
`

type AddBookParams struct {
	Bookname string
	Writer   string
	UserID   pgtype.Int4
	Deadline pgtype.Date
}

func (q *Queries) AddBook(ctx context.Context, arg AddBookParams) error {
	_, err := q.db.Exec(ctx, addBook,
		arg.Bookname,
		arg.Writer,
		arg.UserID,
		arg.Deadline,
	)
	return err
}

const addUser = `-- name: AddUser :exec
INSERT INTO users (email, password, name, surname, roles)
VALUES ($1, $2, $3, $4, $5)
`

type AddUserParams struct {
	Email    string
	Password string
	Name     string
	Surname  string
	Roles    pgtype.Text
}

func (q *Queries) AddUser(ctx context.Context, arg AddUserParams) error {
	_, err := q.db.Exec(ctx, addUser,
		arg.Email,
		arg.Password,
		arg.Name,
		arg.Surname,
		arg.Roles,
	)
	return err
}

const controlUser = `-- name: ControlUser :one
SELECT id, email, password, name, surname, roles FROM users
WHERE email = $1 AND password = $2
`

type ControlUserParams struct {
	Email    string
	Password string
}

func (q *Queries) ControlUser(ctx context.Context, arg ControlUserParams) (User, error) {
	row := q.db.QueryRow(ctx, controlUser, arg.Email, arg.Password)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.Password,
		&i.Name,
		&i.Surname,
		&i.Roles,
	)
	return i, err
}

const deleteBook = `-- name: DeleteBook :exec
DELETE FROM books
WHERE id = $1
`

func (q *Queries) DeleteBook(ctx context.Context, id int32) error {
	_, err := q.db.Exec(ctx, deleteBook, id)
	return err
}

const deleteUser = `-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1
`

func (q *Queries) DeleteUser(ctx context.Context, id int32) error {
	_, err := q.db.Exec(ctx, deleteUser, id)
	return err
}

const listBooks = `-- name: ListBooks :many
SELECT id, bookname, writer, user_id, deadline FROM books
`

func (q *Queries) ListBooks(ctx context.Context) ([]Book, error) {
	rows, err := q.db.Query(ctx, listBooks)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Book
	for rows.Next() {
		var i Book
		if err := rows.Scan(
			&i.ID,
			&i.Bookname,
			&i.Writer,
			&i.UserID,
			&i.Deadline,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listDeadlineBooks = `-- name: ListDeadlineBooks :many
SELECT id, bookname, writer, user_id, deadline FROM books
WHERE deadline < $1
`

func (q *Queries) ListDeadlineBooks(ctx context.Context, deadline pgtype.Date) ([]Book, error) {
	rows, err := q.db.Query(ctx, listDeadlineBooks, deadline)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Book
	for rows.Next() {
		var i Book
		if err := rows.Scan(
			&i.ID,
			&i.Bookname,
			&i.Writer,
			&i.UserID,
			&i.Deadline,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listUsers = `-- name: ListUsers :many
SELECT id, email, password, name, surname, roles FROM users
`

func (q *Queries) ListUsers(ctx context.Context) ([]User, error) {
	rows, err := q.db.Query(ctx, listUsers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []User
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.ID,
			&i.Email,
			&i.Password,
			&i.Name,
			&i.Surname,
			&i.Roles,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateBook = `-- name: UpdateBook :exec
UPDATE books
SET
  bookname = CASE WHEN $2 = '' THEN bookname ELSE $2 END,
  writer = CASE WHEN $3 = '' THEN writer ELSE $3 END,
  user_id = CASE WHEN $4 = 0 THEN user_id ELSE $4 END,
  deadline = CASE WHEN $5 = '' THEN deadline ELSE $5 END
WHERE id = $1
`

type UpdateBookParams struct {
	ID      int32
	Bookname string
	Writer   string
	UserID   pgtype.Int4
	Deadline pgtype.Date
}

func (q *Queries) UpdateBook(ctx context.Context, arg UpdateBookParams) error {
	_, err := q.db.Exec(ctx, updateBook,
		arg.ID,
		arg.Bookname,
		arg.Writer,
		arg.UserID,
		arg.Deadline,
	)
	return err
}

const updateUser = `-- name: UpdateUser :exec
UPDATE users
SET
  email = CASE WHEN $2 = '' THEN email ELSE $2 END,
  password = CASE WHEN $3 = '' THEN password ELSE $3 END,
  name = CASE WHEN $4 = '' THEN name ELSE $4 END,
  surname = CASE WHEN $5 = '' THEN surname ELSE $5 END,
  roles = CASE WHEN $6 = '' THEN roles ELSE $6 END
WHERE id = $1
`

type UpdateUserParams struct {
	ID      int32
	Email string
	Password string
	Name string
	Surname string
	Roles pgtype.Text
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) error {
	_, err := q.db.Exec(ctx, updateUser,
		arg.ID,
		arg.Email,
		arg.Password,
		arg.Name,
		arg.Surname,
		arg.Roles,
	)
	return err
}
