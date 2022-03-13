package repository

import (
	"context"
	"github.com/gocraft/dbr"
	"synergycommunity/internal/domain"
)

type DBConn interface {
	BeginTx(ctx context.Context, f func(tx *dbr.Tx) error) error
}

type Repository struct {
	*UserRepository
	*PostRepository
	*GroupRepository
	*TagRepository
}

type baseConn struct {
	*dbr.Connection
}

func (r *baseConn) BeginTx(ctx context.Context, f func(tx *dbr.Tx) error) error {
	tx, err := r.NewSession(nil).BeginTx(ctx, nil)
	if err != nil {
		return domain.NewDBErrorWrap(err)
	}

	defer tx.RollbackUnlessCommitted()

	err = f(tx)
	if err != nil {
		return domain.NewDBErrorWrap(err)
	}

	if tx.Commit() != nil {
		return domain.NewDBErrorWrap(err)
	}

	return nil
}
func NewRepository(db *dbr.Connection) *Repository {
	base := &baseConn{db}

	return &Repository{
		UserRepository:  NewUserRepository(base),
		PostRepository:  NewPostRepository(base),
		GroupRepository: NewGroupRepository(base),
		TagRepository:   NewTagRepository(base),
	}
}
