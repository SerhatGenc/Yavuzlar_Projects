package controller

import (
	database "library/DB"
	errors "library/Errors"
	models "library/internal/database"

	"github.com/jackc/pgx/v5/pgtype"
)

var (
	db, ctx = database.ConnectDB()
)

func AddUser(email, pass, name, surname string, roles pgtype.Text) {

	err := db.AddUser(ctx, models.AddUserParams{
		Email:    email,
		Password: pass,
		Name:     name,
		Surname:  surname,
		Roles:    roles,
	})
	if err != nil {
		errors.Logger(err, 0)
	}

}

func DeleteUser(id int32) {
	err := db.DeleteUser(ctx, id)

	if err != nil {
		errors.Logger(err, int(id))
	}
}

func UpdateUser(id int32, email, pass, name, surname string, roles pgtype.Text) {

	err := db.UpdateUser(ctx, models.UpdateUserParams{
		ID:       id,
		Email:    email,
		Password: pass,
		Name:     name,
		Surname:  surname,
		Roles:    roles,
	})
	if err != nil {
		errors.Logger(err, 0)
	}
}

func ListUser() []models.User {
	module, err := db.ListUsers(ctx)
	if err != nil {
		errors.Logger(err, 0)
	}

	return module
}

func ControlUser(email, pass string) models.User {
	user, err := db.ControlUser(ctx, models.ControlUserParams{
		Email:    email,
		Password: pass,
	})

	if err != nil {
		errors.Logger(err, 0)
	}

	return user
}
