// Package repository implements structs which contain necessary model methods for different DB entities
package repository

import (
	"context"
	"synergycommunity/internal/domain/entity"
	"synergycommunity/internal/dto"
	"synergycommunity/internal/infrastructure/dbmodel"

	"github.com/gocraft/dbr"
)

type TagRepository struct {
	db DBConn
}

func NewTagRepository(db DBConn) *TagRepository {
	return &TagRepository{db: db}
}

// InsertTag creates a new tag entry in the database and returns its ID.
func (r *TagRepository) InsertTag(ctx context.Context, e entity.Tag) (int64, error) {
	t := dto.TagToDB(e)

	err := r.db.BeginTx(ctx, func(tx *dbr.Tx) error {
		return tx.InsertInto("tags").
			Returning("tag_id").
			Pair("name", t.Name).
			// Pair("created_at", t.CreatedAt).
			// Pair("updated_at", t.UpdatedAt).
			LoadContext(ctx, &t.ID)
	})

	return t.ID, err
}

// Tags returns an array of Tags from the database.
func (r *TagRepository) Tags(ctx context.Context, o entity.Options) ([]entity.Tag, int64, error) {
	var (
		tags  []dbmodel.Tag
		count int64
	)

	err := r.db.BeginTx(ctx, func(tx *dbr.Tx) error {
		stmt := tx.Select("COUNT(tag_id)").
			From("tags")

		stmt = setFilter(stmt, o.Filters)

		_, err := stmt.LoadContext(ctx, &count)
		if err != nil {
			return err
		}

		stmt = tx.Select("*").
			From("tags")
		stmt = setFilter(stmt, o.Filters)
		stmt = setPagination(stmt, o)

		_, err = stmt.LoadContext(ctx, &tags)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, 0, err
	}

	return dto.TagsFromDB(tags), count, nil
}

// TagByID returns a single tag from the database.
func (r *TagRepository) TagByID(ctx context.Context, id int64) (entity.Tag, error) {
	var tag dbmodel.Tag

	err := r.db.BeginTx(ctx, func(tx *dbr.Tx) error {
		return tx.Select("*").From("tags").
			Where("tag_id = ?", id).
			LoadOne(&tag)
	})
	if err != nil {
		return entity.Tag{}, err
	}

	return dto.TagFromDB(tag), nil
}

// UpdateTag updates a single tag row.
func (r *TagRepository) UpdateTag(ctx context.Context, e entity.Tag) (entity.Tag, error) {
	t := dto.TagToDB(e)

	err := r.db.BeginTx(ctx, func(tx *dbr.Tx) error {
		err := tx.Select("id").From("tags").Where("tag_id =?", t.ID).LoadOne(&t.ID)
		if err != nil {
			return err
		}

		err = tx.Update("tags"). // TODO: add check duplicate key
						Returning("created_at").
						Set("nane", t.Name).
						Set("updated_at", t.UpdatedAt).
						Where("tag_id = ?", t.ID).
						Load(&t)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return entity.Tag{}, err
	}

	return dto.TagFromDB(t), err
}

// DeleteTag removes a single tag row.
func (r *TagRepository) DeleteTag(ctx context.Context, id int64) (int64, error) {
	var rowsDeleted int64

	err := r.db.BeginTx(ctx, func(tx *dbr.Tx) error {
		err := tx.Select("*").From("tags").Where("tag_id = ?", id).LoadOne(&id)
		if err != nil {
			return err
		}

		res, err := tx.DeleteFrom("tags").
			Where("tag_id = ?", id).
			ExecContext(ctx)
		if err != nil {
			return err
		}

		rowsDeleted, err = res.RowsAffected()
		if err != nil {
			return err
		}

		// TODO: add cascade deleting from groups posts, subscriptions
		return nil
	})
	if err != nil {
		return 0, err
	}

	return rowsDeleted, nil // TODO rewrite to return only an error
}

func insertTagsToGroup(tx *dbr.Tx, groupID int64, tagIDs []entity.Tag) error {
	if len(tagIDs) == 0 {
		return nil
	}

	stmt, err := tx.Prepare("INSERT INTO groups_tags (group_id, tag_id) VALUES ($1,$2)")
	if err != nil {
		return err
	}

	defer stmt.Close()

	for _, tag := range tagIDs {
		_, err := stmt.Exec(groupID, tag.ID)
		if err != nil {
			return err
		}
	}

	return nil
}

func insertTagsToPost(tx *dbr.Tx, postID int64, tags []dbmodel.Tag) error {
	stmt, err := tx.Prepare("INSERT INTO posts_tags (post_id, tag_id) VALUES ($1,$2)")
	if err != nil {
		return err
	}

	defer stmt.Close()

	for _, tag := range tags {
		_, err := stmt.Exec(postID, tag.ID)
		if err != nil {
			return err
		}
	}

	return nil
}

func updateGroupTags(tx *dbr.Tx, groupID int64, tags []entity.Tag) error {
	_, err := tx.DeleteFrom("groups_tags").
		Where("group_id = ?", groupID).
		Exec()
	if err != nil {
		return err
	}

	p, err := tx.Prepare("INSERT INTO groups_tags (tag_id, group_id) VALUES ($1, $2)")
	if err != nil {
		return err
	}

	defer p.Close()

	for _, t := range tags {
		_, err := p.Exec(t.ID, groupID)
		if err != nil {
			return err
		}
	}

	return nil
}

func updatePostTags(tx *dbr.Tx, postID int64, tags []dbmodel.Tag) error {
	_, err := tx.DeleteFrom("posts_tags").
		Where("post_id = ?", postID).
		Exec()
	if err != nil {
		return err
	}

	p, err := tx.Prepare("INSERT INTO posts_tags (tag_id, post_id) VALUES ($1, $2)")
	if err != nil {
		return err
	}

	defer p.Close()

	for _, t := range tags {
		_, err := p.Exec(t.ID, postID)
		if err != nil {
			return err
		}
	}

	return nil
}

func selectTagsByGroupID(tx *dbr.Tx, id int64) ([]entity.Tag, error) {
	var tags []entity.Tag

	_, err := tx.Select("tags.*").
		From("tags").
		Join("groups_tags", "groups_tags.tag_id = tags.tag_id").
		Where("groups_tags.group_id = ?", id).
		Load(&tags)
	if err != nil {
		return nil, err
	}

	return tags, nil
}
