package userrepository

import (
	"context"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type UserRepository struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (us *UserRepository) Create(ctx context.Context, u *User) error {
	const q = `
		INSERT INTO usrs (id, fio, post, department) 
			VALUES (:id, :fio, :post, :department)
	`
	_, err := us.db.NamedExecContext(ctx, q, u)
	return err
}

func (us *UserRepository) Get(ctx context.Context, id uuid.UUID) (*User, error) {
	const q = `
		SELECT id, fio, post, department FROM usrs WHERE id = $1
	`
	u := new(User)
	err := us.db.GetContext(ctx, u, q, id)
	return u, err
}

func (us *UserRepository) List(ctx context.Context) ([]*User, error) {
	const q = `
		SELECT id, fio, post, department FROM usrs 
	`
	var list []*User
	err := us.db.SelectContext(ctx, &list, q)
	return list, err
}

func (us *UserRepository) Update(ctx context.Context, u *User) error {
	const q = `
		UPDATE usrs SET fio=:fio, post=:post, department=:department 
			WHERE id = :id
	`
	_, err := us.db.NamedExecContext(ctx, q, u)
	return err
}

func (us *UserRepository) Delete(ctx context.Context, id uuid.UUID) error {
	const q = `
		DELETE FROM usrs WHERE id = $1
	`
	_, err := us.db.ExecContext(ctx, q, id)
	return err
}
