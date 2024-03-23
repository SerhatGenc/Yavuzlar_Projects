// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package database

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type Book struct {
	ID       int32
	Bookname string
	Writer   string
	UserID   pgtype.Int4
	Deadline pgtype.Date
}

type User struct {
	ID       int32
	Email    string
	Password string
	Name     string
	Surname  string
	Roles    pgtype.Text
}