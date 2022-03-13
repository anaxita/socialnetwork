package interactor

import (
	"context"
	"synergycommunity/internal/domain/entity"
	"synergycommunity/internal/domain/service"
)

type TagInteractor struct {
	tag *service.TagService
}

func NewTagInteractor(tag *service.TagService) *TagInteractor {
	return &TagInteractor{tag: tag}
}

// Create creates a new tag.
func (i *TagInteractor) Create(ctx context.Context, e entity.Tag) (entity.Tag, error) {
	return i.tag.AddTag(ctx, e)
}

// TagByID returns a tag object with the given ID.
func (i *TagInteractor) TagByID(ctx context.Context, id int64) (entity.Tag, error) {
	return i.tag.ByID(ctx, id)
}

// Tags returns all existing Tags.
func (i *TagInteractor) Tags(ctx context.Context, o entity.Options) ([]entity.Tag, int64, error) {
	return i.tag.Tags(ctx, o)
}

// Edit modifies the tag.
func (i *TagInteractor) Edit(ctx context.Context, e entity.Tag) (entity.Tag, error) {
	return i.tag.Edit(ctx, e)
}

// Delete removes the tag.
func (i *TagInteractor) Delete(ctx context.Context, id int64) error {
	return i.tag.Delete(ctx, id)
}
