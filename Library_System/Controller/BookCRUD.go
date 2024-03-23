package controller

import (
	errors "library/Errors"
	models "library/internal/database"

	"github.com/jackc/pgx/v5/pgtype"
)

func AddBook(bookname, writer string, userId pgtype.Int4, deadline pgtype.Date) {
	err := db.AddBook(ctx, models.AddBookParams{
		Bookname: bookname,
		Writer:   writer,
		UserID:   userId,
		Deadline: deadline,
	})

	if err != nil {
		errors.Logger(err, 0)
	}
}

func DeleteBook(id int32) {
	err := db.DeleteBook(ctx, id)
	if err != nil {
		errors.Logger(err, 0)
	}
}

func UpdateBook(id int32, bookname, writer string, userId pgtype.Int4, deadline pgtype.Date) {
	err := db.UpdateBook(ctx, models.UpdateBookParams{
		ID:       id,
		Bookname: bookname,
		Writer:   writer,
		UserID:   userId,
		Deadline: deadline,
	})
	if err != nil {
		errors.Logger(err, 0)
	}
}

func ListBooks() []models.Book {
	books, err := db.ListBooks(ctx)
	if err != nil {
		errors.Logger(err, 0)
	}
	return books
}

func DeadlineBooks(date pgtype.Date) []models.Book {
	deadBooks, err := db.ListDeadlineBooks(ctx, date)
	if err != nil {
		errors.Logger(err, 0)
	}
	return deadBooks
}
