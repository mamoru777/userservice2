package userrepository

import (
	"context"
	"github.com/google/uuid"
)

type IUserRepository interface {
	Create(ctx context.Context, u *User) error
	Get(ctx context.Context, id uuid.UUID) (*User, error)
	List(ctx context.Context) ([]*User, error)
	Update(ctx context.Context, u *User) error
	Delete(ctx context.Context, id uuid.UUID) error
}
