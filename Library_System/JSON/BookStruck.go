package json

import "github.com/jackc/pgx/v5/pgtype"

type AddBookStr struct {
	Bookname string      `json:"bookname"`
	Writer   string      `json:"writer"`
	UserID   pgtype.Int4 `json:"userId"`
	Deadline pgtype.Date `json:"deadline"`
}

type DeleteBookStr struct {
	ID int32 `json:"id"`
}

type FullBookStr struct {
	ID       int32       `json:"id"`
	Bookname string      `json:"bookname"`
	Writer   string      `json:"writer"`
	UserID   pgtype.Int4 `json:"userId"`
	Deadline pgtype.Date `json:"deadline"`
}
