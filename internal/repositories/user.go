package repositories

import (
	"api-dev/domain"
	"context"
	"database/sql"
	"time"

	"github.com/doug-martin/goqu/v9"
)

type userRepository struct {
	db *goqu.Database
}

func NewUser(con *goqu.Database) domain.UserRepository {
	return &userRepository{
		db: con,
	}
}

// FindAll implements domain.UserRepository.
func (ur *userRepository) FindAll(ctx context.Context) (result []domain.User, err error) {
	dataset := ur.db.From("users").Where(goqu.C("deleted_at").IsNull())
	err = dataset.ScanStructsContext(ctx, &result)
	return
}

// FindById implements domain.UserRepository.
func (ur *userRepository) FindById(ctx context.Context, id string) (result domain.User, err error) {
	dataset := ur.db.From("users").Where(goqu.C("deleted_at").IsNull(), goqu.C("id").Eq(id))
	_, err = dataset.ScanStructContext(ctx, &result)
	return
}

// Save implements domain.UserRepository.
func (ur *userRepository) Save(ctx context.Context, u *domain.User) error {
	executor := ur.db.Insert("users").Rows(u).Executor()
	_, err := executor.ExecContext(ctx)
	return err
}

// Update implements domain.UserRepository.
func (ur *userRepository) Update(ctx context.Context, u *domain.User) error {
	executor := ur.db.Update("users").Where(goqu.C("id").Eq(u.Id)).Set(u).Executor()
	_, err := executor.ExecContext(ctx)
	return err
}

// Delete implements domain.UserRepository.
func (ur *userRepository) Delete(ctx context.Context, id string) error {
	executor := ur.db.Update("users").Where(goqu.C("id").Eq(id)).Set(goqu.Record{"deleted_at": sql.NullTime{Valid: true, Time: time.Now()}}).Executor()
	_, err := executor.ExecContext(ctx)
	return err
}
