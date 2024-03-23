package json

import "github.com/jackc/pgx/v5/pgtype"

type AddUserStr struct {
	Email    string      `json:"email"`
	Password string      `json:"password"`
	Name     string      `json:"name"`
	Surname  string      `json:"surname"`
	Roles    pgtype.Text `json:"roles"`
}
type DeleteUserStr struct {
	ID int32 `json:"id"`
}
type FullUserStr struct {
	ID       int32       `json:"id"`
	Email    string      `json:"email"`
	Password string      `json:"password"`
	Name     string      `json:"name"`
	Surname  string      `json:"surname"`
	Roles    pgtype.Text `json:"roles"`
}

type ControlUserStr struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
