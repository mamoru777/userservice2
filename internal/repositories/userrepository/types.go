package userrepository

import "github.com/google/uuid"

type User struct {
	Id         uuid.UUID `db:"id"`
	Fio        string    `db:"fio"`
	Post       string    `db:"post"`
	Department string    `db:"department"`
}
