package repository

import (
	"context"
	"synergycommunity/internal/domain"
	"synergycommunity/internal/domain/entity"
	"synergycommunity/internal/dto"
	"synergycommunity/internal/infrastructure/dbmodel"
	"time"

	"github.com/gocraft/dbr"
)

type PostRepository struct {
	db DBConn
}

func NewPostRepository(db DBConn) *PostRepository {
	return &PostRepository{db: db}
}

// InsertPost creates a new Post entry in the database and returns its ID.
func (r *PostRepository) InsertPost(ctx context.Context, e entity.Post) (int64, error) {
	p, tags := dto.PostToDB(e)

	err := r.db.BeginTx(ctx, func(tx *dbr.Tx) error { // TODO add check existing group
		err := tx.InsertInto("posts").
			Returning("post_id").
			Columns("user_id", "group_id", "name", "text", "created_at", "updated_at").
			Record(&p).
			LoadContext(ctx, &p.ID)
		if err != nil {
			return err
		}

		err = insertTagsToPost(tx, p.ID, tags)
		if err != nil {
			return err
		}

		return nil
	})

	return p.ID, err
}

// Posts returns an array of Posts from the database.
func (r *PostRepository) Posts(ctx context.Context, opts entity.Options) ([]entity.Post, int64, error) {
	var (
		posts []dbmodel.Post
		count int64
	)

	err := r.db.BeginTx(ctx, func(tx *dbr.Tx) error {
		stmt := r.selectPosts(tx, "COUNT(post_id)")
		stmt = setFilter(stmt, opts.Filters)

		_, err := stmt.LoadContext(ctx, &count)
		if err != nil {
			return err
		}

		stmt = r.selectPostsWithTags(tx)
		stmt = setFilter(stmt, opts.Filters)
		stmt = setPagination(stmt, opts)

		_, err = stmt.LoadContext(ctx, &posts)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, 0, err
	}

	return dto.PostsFromDB(posts), count, nil
}

// PostsBySubscriptions returns an array of Posts from the database.
func (r *PostRepository) PostsBySubscriptions(ctx context.Context, opts entity.Options) ([]entity.Post, int64, error) {
	var (
		posts []dbmodel.Post
		count int64
	)

	err := r.db.BeginTx(ctx, func(tx *dbr.Tx) error {
		_, err := tx.SelectBySql(`
	SELECT posts.*, JSON_AGG(t) tags
	FROM posts
         JOIN users ON users.user_id = 2
         JOIN groups_subs ON users.user_id = groups_subs.user_id
         JOIN users_subs us ON users.user_id = us.user_id
         JOIN tags_subs ts ON users.user_id = ts.user_id
         JOIN groups g ON g.group_id = groups_subs.group_id
         JOIN posts_tags pt ON ts.tag_id = pt.tag_id
         JOIN tags t ON t.tag_id = ts.tag_id
	WHERE posts.group_id = g.group_id
   		OR posts.post_id = pt.post_id
   		OR posts.user_id = us.to_user_id
	GROUP BY posts.post_id
	ORDER BY posts.post_id`).Load(&posts)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, 0, err
	}

	return dto.PostsFromDB(posts), count, nil
}

// PostByID returns a single Post from the database.
func (r *PostRepository) PostByID(ctx context.Context, id int64) (entity.Post, error) {
	var (
		post entity.Post
		tags []entity.Tag
		err  error
	)

	err = r.db.BeginTx(ctx, func(tx *dbr.Tx) error {
		post, err = selectPostByID(tx, id, "*")
		if err != nil {
			return err
		}

		tags, err = r.selectTagsByPostID(tx, id)
		if err != nil {
			return err
		}

		post.Tags = tags

		return nil
	})
	if err != nil {
		return entity.Post{}, err
	}

	return post, nil
}

// UpdatePost updates a single Post row.
func (r *PostRepository) UpdatePost(ctx context.Context, e entity.Post) error {
	p, tags := dto.PostToDB(e)

	err := r.db.BeginTx(ctx, func(tx *dbr.Tx) error { // TODO add check existing post
		_, err := tx.Update("posts").
			Returning("post_id", "created_at").
			Set("name", p.Name).
			Set("text", p.Text).
			Set("updated_at", p.UpdatedAt).
			Where("post_id = ?", p.ID).
			Where("deleted_at IS NULL").
			Exec()
		if err != nil {
			return err
		}

		err = updatePostTags(tx, p.ID, tags)
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

// DeletePost removes a single Post row.
func (r *PostRepository) DeletePost(ctx context.Context, id int64) error {
	return r.db.BeginTx(ctx, func(tx *dbr.Tx) error {
		_, err := selectPostByID(tx, id, "id")
		if err != nil {
			return err
		}

		_, err = tx.
			Update("posts").
			Where("post_id = ?", id).
			Set("deleted_at", time.Now().UTC()).
			ExecContext(ctx)
		if err != nil {
			return err
		}

		return nil
	})
}

func (r *PostRepository) selectPosts(tx *dbr.Tx, cols ...string) *dbr.SelectStmt {
	return tx.Select(cols...).
		From("posts").
		Where("deleted_at IS NULL")
}

func selectPostByID(tx *dbr.Tx, postID int64, cols ...string) (entity.Post, error) {
	var p entity.Post
	err := tx.Select(cols...).
		From("posts").
		Where("deleted_at IS NULL").
		Where("post_id = ?", postID).LoadOne(&p)

	return p, err
}

func (r *PostRepository) selectPostsWithTags(tx *dbr.Tx) *dbr.SelectStmt {
	return r.selectPosts(tx, "posts.*", "JSONB_AGG(tags.*) FILTER ( WHERE tags.tag_id IS NOT NULL ) tags").
		LeftJoin("posts_tags", "posts.post_id = posts_tags.post_id").
		LeftJoin("tags", "posts_tags.tag_id = tags.tag_id").
		Where("posts.deleted_at IS NULL").
		GroupBy("posts.post_id")
}

func (r *PostRepository) selectTagsByPostID(tx *dbr.Tx, postID int64) ([]entity.Tag, error) {
	var tags []entity.Tag

	_, err := tx.Select("*").
		From("tags").
		Join("posts_tags", "tags.tag_id = posts_tags.tag_id").
		Where("posts_tags.post_id = ?", postID).
		Load(&tags)
	if err != nil {
		return nil, domain.NewDBErrorWrap(err)
	}

	return tags, nil
}

func (r *PostRepository) SelectUserPermissionsByPostID(ctx context.Context, userID int64, postID int64) ([]domain.Perm, error) {
	var (
		permissions []domain.Perm
		err         error
	)

	err = r.db.BeginTx(ctx, func(tx *dbr.Tx) error {
		p, err := selectPostByID(tx, postID, "group_id")
		if err != nil {
			return err
		}

		permissions, err = selectUserPermissionsByGroupID(tx, userID, p.GroupID)

		return err
	})

	return permissions, err
}
