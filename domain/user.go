package domain

import (
	"api-dev/dto"
	"context"
	"database/sql"
)

type User struct {
	Id        int          `db:"id"`
	Name      string       `db:"name"`
	Email     string       `db:"email"`
	Password  string       `db:"password"`
	CreatedAt sql.NullTime `db:"created_at"`
	UpdatedAt sql.NullTime `db:"updated_at"`
	DeletedAt sql.NullTime `db:"deleted_at"`
}

type UserRepository interface {
	FindAll(ctx context.Context) ([]User, error)
	FindById(ctx context.Context, id string) (User, error)
	Save(ctx context.Context, u *User) error
	Update(ctx context.Context, u *User) error
	Delete(ctx context.Context, id string) error
}

type UserService interface {
	Index(ctx context.Context) ([]dto.UserData, error)
}
