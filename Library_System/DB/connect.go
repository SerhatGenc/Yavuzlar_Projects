package db

import (
	"context"
	errors "library/Errors"
	db "library/internal/database"

	"github.com/jackc/pgx/v5"
)

func ConnectDB() (*db.Queries, context.Context) {
	ctx := context.Background()

	conn, err := pgx.Connect(ctx, "user=root password=root host=localhost port=5432 dbname=library sslmode=disable")
	if err != nil {
		errors.Logger(err, 0)
	}

	queries := db.New(conn)

	return queries, ctx
}
