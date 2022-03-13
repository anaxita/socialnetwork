package service

import (
	"context"
	"errors"
	"synergycommunity/internal/domain"
	"synergycommunity/internal/domain/entity"
	"synergycommunity/internal/dto"
	"time"

	"github.com/gocraft/dbr"
)

type PostStorage interface {
	InsertPost(ctx context.Context, e entity.Post) (int64, error)
	Posts(ctx context.Context, opts entity.Options) ([]entity.Post, int64, error)
	PostByID(ctx context.Context, id int64) (entity.Post, error)
	UpdatePost(ctx context.Context, e entity.Post) error
	DeletePost(ctx context.Context, id int64) error
	SelectUserPermissionsByPostID(ctx context.Context, userID int64, postID int64) ([]domain.Perm, error)
}

type PostService struct {
	repo PostStorage
}

func NewPostService(post PostStorage) *PostService {
	return &PostService{
		repo: post,
	}
}

// AddPost creates a new Post.
func (s *PostService) AddPost(ctx context.Context, e entity.Post) (entity.Post, error) {
	session := dto.Session(ctx)

	e.CreatedAt = time.Now().UTC()
	e.UpdatedAt = e.CreatedAt
	e.UserID = session.User.ID

	id, err := s.repo.InsertPost(ctx, e)
	if err != nil {
		return entity.Post{}, err
	}

	e.ID = id

	return e, nil
}

// ByID returns a Post by its ID.
func (s *PostService) ByID(ctx context.Context, id int64) (entity.Post, error) {
	post, err := s.repo.PostByID(ctx, id)
	if err != nil {
		if errors.Is(err, dbr.ErrNotFound) {
			return entity.Post{}, nil
		}

		return entity.Post{}, err
	}

	return post, nil
}

// Posts returns all Post objects.
func (s *PostService) Posts(ctx context.Context, opts entity.Options) (
	[]entity.Post, int64, error,
) {
	posts, count, err := s.repo.Posts(ctx, opts)
	if err != nil {
		if errors.Is(err, dbr.ErrNotFound) {
			return make([]entity.Post, 0), 0, nil
		}

		return nil, 0, err
	}

	return posts, count, nil
}

// PostsByUserSubscriptions ...
func (s *PostService) PostsByUserSubscriptions(ctx context.Context, opts entity.Options) (
	[]entity.Post, int64, error,
) {
	posts, count, err := s.repo.Posts(ctx, opts)
	if err != nil {
		if errors.Is(err, dbr.ErrNotFound) {
			return make([]entity.Post, 0), 0, nil
		}

		return nil, 0, err
	}

	return posts, count, nil
}

// Edit modifies a Post.
func (s *PostService) Edit(ctx context.Context, e entity.Post) (entity.Post, error) {
	e.UpdatedAt = time.Now().UTC()

	err := s.repo.UpdatePost(ctx, e)
	if err != nil {
		return entity.Post{}, err
	}

	return e, nil
}

// Delete removes a Post.
func (s *PostService) Delete(ctx context.Context, id int64) error {
	err := s.repo.DeletePost(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

func (s *PostService) UserPermissions(ctx context.Context, userID int64, postID int64) ([]domain.Perm, error) {
	return s.repo.SelectUserPermissionsByPostID(ctx, userID, postID)
}
