// Package repository implements structs which contain necessary model methods for different DB entities
package repository

import (
	"context"
	"errors"
	"github.com/gocraft/dbr"
	"synergycommunity/internal/domain"
	"synergycommunity/internal/domain/entity"
	"synergycommunity/internal/dto"
	"synergycommunity/internal/infrastructure/dbmodel"
	"time"
)

type UserRepository struct {
	db DBConn
}

func NewUserRepository(db DBConn) *UserRepository {
	return &UserRepository{db: db}
}

// InsertUser creates a new User entry in the database and returns its ID.
func (r *UserRepository) InsertUser(ctx context.Context, e entity.User) (int64, error) {
	u := dto.UserToDB(e)

	err := r.db.BeginTx(ctx, func(tx *dbr.Tx) error {
		return tx.InsertInto("users").
			Returning("user_id").
			Columns("nickname", "first_name", "last_name", "created_at", "updated_at").
			Record(&u).
			Load(&u.ID)
	})
	if err != nil {
		return 0, err
	}

	return u.ID, nil

}

// Users returns an array of Users from the database.
func (r *UserRepository) Users(ctx context.Context, opts entity.Options) ([]entity.User, int64,
	error) {
	var (
		users []dbmodel.User
		count int64
	)

	err := r.db.BeginTx(ctx, func(tx *dbr.Tx) error {
		stmt := tx.
			Select("COUNT(user_id)").
			From("users")

		stmt = setFilter(stmt, opts.Filters)

		_, err := stmt.LoadContext(ctx, &count)
		if err != nil {
			return err
		}

		stmt = selectUsers(tx, "*")
		stmt = setFilter(stmt, opts.Filters)
		stmt = setPagination(stmt, opts)

		_, err = stmt.LoadContext(ctx, &users)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, 0, err
	}

	return dto.UsersFromDB(users), count, nil
}

// UserByID returns a single User from the database.
func (r *UserRepository) UserByID(ctx context.Context, id int64) (entity.User, error) {
	var u dbmodel.User

	err := r.db.BeginTx(ctx, func(tx *dbr.Tx) error {
		err := selectUsers(tx, "*").
			Where("user_id = ?", id).
			LoadOne(&u)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return entity.User{}, err
	}

	return dto.UserFromDB(u), nil
}

// UpdateUser updates a single User row.
func (r *UserRepository) UpdateUser(ctx context.Context, e entity.User) (entity.User, error) {
	u := dto.UserToDB(e)

	err := r.db.BeginTx(ctx, func(tx *dbr.Tx) error {
		return tx.Update("users").
			Returning("created_at").
			Set("nickname", u.Nickname).
			Set("first_name", u.FirstName).
			Set("last_name", u.LastName).
			Set("updated_at", u.UpdatedAt).
			Where("user_id = ?", u.ID).
			Where("deleted_at IS NULL").
			Load(&u)
	})
	if err != nil {
		return entity.User{}, err
	}

	if u.CreatedAt.IsZero() {
		return entity.User{}, domain.NewDBErrorWrap(dbr.ErrNotFound)
	}

	return dto.UserFromDB(u), nil
}

// DeleteUser removes a single User row.
func (r *UserRepository) DeleteUser(ctx context.Context, id int64) error {
	return r.db.BeginTx(ctx, func(tx *dbr.Tx) error {
		err := selectUsers(tx, "user_id").Where("id =?", id).LoadOne(&id)
		if err != nil {
			return err
		}

		_, err = tx.Update("groups").
			Set("deleted_at", time.Now().UTC()).
			Where("user_id = ?", id).
			ExecContext(ctx)

		return err
	})
}

func (r *UserRepository) InsertSubscription(ctx context.Context, e entity.Subscription) error {
	var (
		table  string
		column string
	)

	switch e.Model {
	case entity.ModelGroup:
		table = "groups_subs"
		column = "group_id"
	case entity.ModelUser:
		table = "users_subs"
		column = "to_user_id"
	case entity.ModelTag:
		table = "tags_subs"
		column = "tag_id"
	}

	return r.db.BeginTx(ctx, func(tx *dbr.Tx) error {
		_, err := tx.InsertInto(table).
			Pair("user_id", e.UserID).
			Pair(column, e.ModelID).
			Exec()
		return err
	})
}

func (r *UserRepository) DeleteSubscription(ctx context.Context, e entity.Subscription) error {
	var (
		table  string
		column string
	)

	switch e.Model {
	case entity.ModelGroup:
		table = "groups_subs"
		column = "group_id"
	case entity.ModelUser:
		table = "users_subs"
		column = "to_user_id"
	case entity.ModelTag:
		table = "tags_subs"
		column = "tag_id"
	}

	return r.db.BeginTx(ctx, func(tx *dbr.Tx) error {
		_, err := tx.DeleteFrom(table).
			Where("user_id = ?", e.UserID).
			Where(dbr.Eq(column, e.ModelID)).
			Exec()

		return err
	})
}

func (r *UserRepository) Subscriptions(ctx context.Context, userID int64) (entity.UserSubscriptions, error) {
	var subs entity.UserSubscriptions

	err := r.db.BeginTx(ctx, func(tx *dbr.Tx) error {
		_, err := tx.Select("groups.*").
			From("groups").
			Join("groups_subs",
				dbr.And(
					dbr.Eq("groups_subs.user_id", userID),
					dbr.Eq("groups_subs.group_id", dbr.I("groups.id"))),
			).
			Load(&subs.Groups)
		if err != nil && !errors.Is(err, dbr.ErrNotFound) {
			return err
		}

		_, err = tx.Select("tags.*").
			From("tags").
			Join("tags_subs",
				dbr.And(
					dbr.Eq("tags_subs.user_id", userID),
					dbr.Eq("tags_subs.tag_id", dbr.I("tags.id"))),
			).
			Load(&subs.Tags)
		if err != nil && !errors.Is(err, dbr.ErrNotFound) {
			return err
		}

		_, err = tx.Select("users.*").
			From("users").
			Join("users_subs",
				dbr.And(
					dbr.Eq("users_subs.user_id", userID),
					dbr.Eq("users_subs.to_user_id", dbr.I("users.id"))),
			).
			Load(&subs.Users)
		if err != nil && !errors.Is(err, dbr.ErrNotFound) {
			return err
		}

		return nil
	})

	if err != nil {
		return entity.UserSubscriptions{}, err
	}

	return subs, nil
}

func selectUsers(tx *dbr.Tx, cols ...string) *dbr.SelectStmt {
	return tx.
		Select(cols...).
		From("users").
		Where("deleted_at IS NULL")
}
