// Package repository implements structs which contain necessary model methods for different DB entities
package repository

import (
	"context"
	"errors"
	"synergycommunity/internal/domain"
	"synergycommunity/internal/domain/entity"
	"synergycommunity/internal/dto"
	"synergycommunity/internal/infrastructure/dbmodel"
	"time"

	"github.com/gocraft/dbr"
)

type GroupRepository struct {
	db DBConn
}

func NewGroupRepository(db DBConn) *GroupRepository {
	return &GroupRepository{db: db}
}

// InsertGroup creates a new Group entry in the database and returns its ID.
func (r *GroupRepository) InsertGroup(ctx context.Context, e entity.Group) (int64, error) {
	err := r.db.BeginTx(ctx, func(tx *dbr.Tx) error {
		_, err := selectGroups(tx, "group_id").Where("slug = ?", e.Slug).Load(&e.ID)
		if err != nil && !errors.Is(err, dbr.ErrNotFound) {
			return err
		}

		if e.ID != 0 {
			return domain.ErrDuplicateKey // TODO система ошибок говно, переделай
		}

		err = tx.InsertInto("groups").
			Returning("group_id").
			Columns("user_id", "name", "slug", "description", "created_at", "updated_at").
			Record(&e).
			Load(&e.ID)
		if err != nil {
			return err
		}

		err = insertTagsToGroup(tx, e.ID, e.Tags)
		if err != nil {
			return err
		}

		ur := dbmodel.UserRole{
			UserID:  e.UserID,
			GroupID: e.ID,
			RoleID:  int64(domain.RoleAdmins),
		}

		err = insertUserRole(tx, ur)
		if err != nil {
			return err
		}

		return nil
	})

	return e.ID, err
}

// Groups returns an array of Groups from the database.
func (r *GroupRepository) Groups(ctx context.Context, opts entity.Options) ([]entity.Group, int64, error) {
	var (
		groups []dbmodel.GroupWithTags
		count  int64
	)

	err := r.db.BeginTx(ctx, func(tx *dbr.Tx) error {
		stmt := selectGroups(tx, "COUNT(group_id)")
		stmt = setFilter(stmt, opts.Filters)

		err := stmt.LoadOne(&count)
		if err != nil {
			return err
		}

		stmt = r.selectGroupsWithTags(tx)
		stmt = setFilter(stmt, opts.Filters)
		stmt = setPagination(stmt, opts)

		_, err = stmt.Load(&groups)
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, 0, err
	}

	return dto.GroupsFromDB(groups), count, nil
}

// SelectByID returns a single Group from the database.
func (r *GroupRepository) SelectByID(ctx context.Context, id int64) (entity.Group, error) {
	var (
		group dbmodel.GroupWithTags
	)

	err := r.db.BeginTx(ctx, func(tx *dbr.Tx) error {
		err := r.selectGroupsWithTags(tx).
			Where("group_id = ?", id).
			LoadOne(&group)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return entity.Group{}, err
	}

	return dto.GroupFromDB(group), nil
}

func selectGroups(tx *dbr.Tx, cols ...string) *dbr.SelectStmt {
	return tx.Select(cols...).
		From("groups").
		Where("deleted_at IS NULL")
}

func selectGroupByID(tx *dbr.Tx, id int64, cols ...string) (entity.Group, error) {
	var g entity.Group
	err := tx.Select(cols...).
		From("groups").
		Where("deleted_at IS NULL").
		Where("group_id = ?", id).
		LoadOne(&g)
	if err != nil {
		return g, err
	}

	return g, nil
}

// UpdateGroup updates a single Group row.
func (r *GroupRepository) UpdateGroup(ctx context.Context, e entity.Group) (entity.Group, error) {

	err := r.db.BeginTx(ctx, func(tx *dbr.Tx) error {
		// TODO: check with DB
		_, err := selectGroupByID(tx, e.ID, "group_id")
		if err != nil {
			return err
		}

		err = tx.Update("groups").
			Returning("created_at").
			Set("slug", e.Slug).
			Set("name", e.Name).
			Set("description", e.Description).
			Set("updated_at", e.UpdatedAt).
			Where("group_id = ?", e.ID).
			Load(&e)
		if err != nil {
			return err
		}

		err = updateGroupTags(tx, e.ID, e.Tags)
		if err != nil {
			return err
		}

		e.Tags, err = selectTagsByGroupID(tx, e.ID)

		return err
	})
	if err != nil {
		return entity.Group{}, err
	}

	return e, nil
}

// DeleteGroup removes a single Group row.
func (r *GroupRepository) DeleteGroup(ctx context.Context, id int64) error {
	return r.db.BeginTx(ctx, func(tx *dbr.Tx) error {
		// TODO: check by DB
		_, err := selectGroupByID(tx, id, "id")
		if err != nil {
			return err
		}

		_, err = tx.Update("groups").
			Set("deleted_at", time.Now().UTC()).
			Where("group_id = ?", id).
			ExecContext(ctx)

		return err
	})
}

func (r *GroupRepository) selectGroupsWithTags(tx *dbr.Tx) *dbr.SelectStmt {
	return selectGroups(tx, "groups.*", "JSON_AGG(tags.*) tags").
		Join("groups_tags", "groups.group_id = groups_tags.group_id").
		Join("tags", "groups_tags.tag_id = tags.tag_id").
		GroupBy("groups.group_id")
}

func (r *GroupRepository) SelectUserPermissions(ctx context.Context, userID int64, groupIDs ...int64) ([]domain.Perm, error) {
	var (
		permissions []domain.Perm
		err         error
	)

	err = r.db.BeginTx(ctx, func(tx *dbr.Tx) error {
		permissions, err = selectUserPermissionsByGroupID(tx, userID, groupIDs...)

		return err
	})

	return permissions, err
}

func selectUserPermissionsByGroupID(tx *dbr.Tx, userID int64, groupIDs ...int64) ([]domain.Perm, error) {
	var permissions []domain.Perm

	_, err := tx.Select("roles_permissions.permission_id").
		From("roles_permissions").
		Join("groups_users_roles", "groups_users_roles.role_id = roles_permissions.role_id").
		Where("user_id = ?", userID).
		Where(
			dbr.Or(
				dbr.Eq("group_id", nil),
				dbr.Eq("group_id", groupIDs),
			),
		).
		Where(
			dbr.Or(
				dbr.Eq("expires_at", nil),
				dbr.Gt("expires_at", time.Now().UTC()),
			),
		).
		GroupBy("roles_permissions.permission_id").
		Load(&permissions)
	if err != nil {
		return nil, err
	}

	return permissions, nil
}
func (r *GroupRepository) InsertUserRole(ctx context.Context, e entity.UserRole) error {
	ur := dto.UserRoleToDB(e)

	return r.db.BeginTx(ctx, func(tx *dbr.Tx) error {
		err := r.hasMember(tx, e.UserID, e.GroupID)
		if err != nil {
			return err
		}

		err = insertUserRole(tx, ur)
		if err != nil {
			return err
		}

		return nil
	})
}
func insertUserRole(tx *dbr.Tx, ur dbmodel.UserRole) error {
	_, err := tx.InsertInto("groups_users_roles").
		Columns("user_id", "group_id", "role_id", "expires_at").
		Record(ur).
		Exec()
	if err != nil {
		return err
	}

	return nil
}

func (r *GroupRepository) DeleteUserRole(ctx context.Context, e entity.UserRole) error {
	err := r.db.BeginTx(ctx, func(tx *dbr.Tx) error {
		err := r.hasMember(tx, e.UserID, e.GroupID)
		if err != nil {
			return err
		}

		_, err = tx.DeleteFrom("group_roles").
			Where("user_id", e.UserID).
			Where("group_id", e.GroupID).
			Where("role_id", e.RoleID).
			Exec()
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return err
	}

	return nil
}

func (r *GroupRepository) hasMember(tx *dbr.Tx, userID, groupID int64) error {
	return tx.Select("user_id").
		From("subscriptions").
		Where("model_id", groupID).
		Where("model", entity.ModelGroup).
		Where("user_id", userID).
		LoadOne(&userID)
}
