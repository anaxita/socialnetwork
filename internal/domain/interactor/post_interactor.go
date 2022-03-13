package interactor

import (
	"context"
	"synergycommunity/internal/domain"
	"synergycommunity/internal/domain/entity"
	"synergycommunity/internal/domain/service"
)

type PostInteractor struct {
	post *service.PostService
}

func NewPostInteractor(ps *service.PostService) *PostInteractor {
	return &PostInteractor{post: ps}
}

// Add creates a new Post.
func (i *PostInteractor) Add(ctx context.Context, e entity.Post) (entity.Post, error) {
	return i.post.AddPost(ctx, e)
}

// PostByID returns a Post object with the given ID.
func (i *PostInteractor) PostByID(ctx context.Context, id int64) (entity.Post, error) {
	return i.post.ByID(ctx, id)
}

// Posts returns all existing Posts.
func (i *PostInteractor) Posts(ctx context.Context, opts entity.Options) ([]entity.Post, int64, error) {
	return i.post.Posts(ctx, opts)
}

// Edit modifies the Post.
func (i *PostInteractor) Edit(ctx context.Context, e entity.Post) (entity.Post, error) {
	return i.post.Edit(ctx, e)
}

// Delete removes the Post.
func (i *PostInteractor) Delete(ctx context.Context, id int64) error {
	return i.post.Delete(ctx, id)
}

// UserPermissions ...
func (i *PostInteractor) UserPermissions(ctx context.Context, userID, postID int64) ([]domain.Perm, error) {
	return i.post.UserPermissions(ctx, userID, postID)
}

// ByUserSubscriptions returns Posts by user subscriptions.
func (i *PostInteractor) ByUserSubscriptions(ctx context.Context, opts entity.Options, userID int64) ([]entity.Post, int64, error) {
	return i.post.PostsByUserSubscriptions(ctx, opts)
}
