// Package service prepares the data to be stored
package service

import (
	"context"
	"errors"
	"synergycommunity/internal/domain"
	"synergycommunity/internal/domain/entity"
	"time"

	"github.com/gocraft/dbr"
)

type TagStorage interface {
	InsertTag(ctx context.Context, e entity.Tag) (int64, error)
	Tags(ctx context.Context, o entity.Options) ([]entity.Tag, int64, error)
	TagByID(ctx context.Context, id int64) (entity.Tag, error)
	UpdateTag(ctx context.Context, e entity.Tag) (entity.Tag, error)
	DeleteTag(ctx context.Context, id int64) (int64, error)
}

type TagService struct {
	repo TagStorage
}

func NewTagService(repo TagStorage) *TagService {
	return &TagService{repo: repo}
}

// AddTag creates a new tag.
func (s *TagService) AddTag(ctx context.Context, e entity.Tag) (entity.Tag, error) {
	e.CreatedAt = time.Now().UTC()
	e.UpdatedAt = e.CreatedAt

	id, err := s.repo.InsertTag(ctx, e)
	if err != nil {
		return entity.Tag{}, err
	}

	e.ID = id

	return e, nil
}

// ByID returns a tag by its ID.
func (s *TagService) ByID(ctx context.Context, id int64) (entity.Tag, error) {
	tag, err := s.repo.TagByID(ctx, id)
	if err != nil {
		return entity.Tag{}, err
	}

	if tag == (entity.Tag{}) {
		return entity.Tag{}, domain.NewErrorWrap(err, domain.ErrCodeNotFound, "tag", "ID", id)
	}

	return tag, nil
}

// Tags returns all tag objects.
func (s *TagService) Tags(ctx context.Context, o entity.Options) ([]entity.Tag, int64, error) {
	tags, count, err := s.repo.Tags(ctx, o)
	if err != nil {
		if errors.Is(err, dbr.ErrNotFound) {
			return make([]entity.Tag, 0), 0, nil
		}

		return nil, 0, err
	}

	return tags, count, nil
}

// Edit modifies an existing tag.
func (s *TagService) Edit(ctx context.Context, e entity.Tag) (entity.Tag, error) {
	e.UpdatedAt = time.Now().UTC()

	tag, err := s.repo.UpdateTag(ctx, e)
	if err != nil {
		return entity.Tag{}, err
	}

	return tag, nil
}

// Delete removes an existing tag.
func (s *TagService) Delete(ctx context.Context, id int64) error {
	rows, err := s.repo.DeleteTag(ctx, id)
	if err != nil {
		return err
	}

	if rows == 0 {
		return domain.NewErrorWrap(err, domain.ErrCodeNotFound, "tag", "ID", id)
	}

	return nil
}
