package gql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"synergycommunity/internal/delivery/api/apimodel"
	"synergycommunity/internal/delivery/api/helpers"
	"synergycommunity/internal/domain"
	"synergycommunity/internal/dto"
)

func (r *mutationResolver) AddPost(ctx context.Context, p apimodel.AddPostInput) (*apimodel.Post, error) {
	s := dto.Session(ctx)
	if !s.IsAuthorized() {
		return nil, domain.NewError(domain.ErrCodeNotAuthorized)
	}

	userPerms, err := r.Groups.UserPermissions(ctx, s.User.ID, p.GroupID)
	if err != nil {
		return nil, err
	}

	disallowed := []domain.Perm{domain.PermNoWriting}
	allowed := []domain.Perm{domain.PermWrite}

	hasAccess := helpers.HasAccess(userPerms, allowed, disallowed)
	if !hasAccess {
		return nil, domain.NewError(domain.ErrCodeForbidden)
	}

	post, err := r.Posts.Add(ctx, p.ToEntity())
	if err != nil {
		return nil, err
	}

	data := dto.PostToRest(post)

	return &data, nil
}

func (r *mutationResolver) EditPost(ctx context.Context, p apimodel.EditPostInput) (*apimodel.Post, error) {
	s := dto.Session(ctx)
	if !s.IsAuthorized() {
		return nil, domain.NewError(domain.ErrCodeNotAuthorized)
	}

	userPerms, err := r.Posts.UserPermissions(ctx, s.User.ID, p.ID)
	if err != nil {
		return nil, err
	}

	ok := helpers.IsAccessAllowed(userPerms, domain.PermEdit, domain.PermWrite)
	if !ok {
		return nil, domain.NewError(domain.ErrCodeForbidden)
	}

	post, err := r.Posts.Edit(ctx, p.ToEntity())
	if err != nil {
		return nil, err
	}

	data := dto.PostToRest(post)

	return &data, nil
}

func (r *mutationResolver) DeletePost(ctx context.Context, id int64) (*bool, error) {
	s := dto.Session(ctx)
	if !s.IsAuthorized() {
		return nil, domain.NewError(domain.ErrCodeNotAuthorized)
	}

	userPerms, err := r.Posts.UserPermissions(ctx, s.User.ID, id)
	if err != nil {
		return nil, err
	}

	isAllowed := helpers.IsAccessAllowed(userPerms, domain.PermDelete)
	if !isAllowed {
		return nil, domain.NewError(domain.ErrCodeForbidden)
	}

	err = r.Posts.Delete(ctx, id)
	if err != nil {
		return nil, err
	}

	data := true

	return &data, nil
}

func (r *queryResolver) ShowPosts(ctx context.Context, o *apimodel.OptionsInput) (*apimodel.PostsWithPagination, error) {
	helpers.ValidateOptions(
		o, apimodel.ColUserID, apimodel.ColGroupID, apimodel.ColText, apimodel.ColName,
	)

	posts, count, err := r.Posts.Posts(ctx, dto.OptionsFromRest(o))
	if err != nil {
		return nil, err
	}

	data := dto.PostsWithPaginationToRest(posts, *o, count)

	return &data, nil
}

func (r *queryResolver) ShowSubscribedPosts(ctx context.Context, o *apimodel.OptionsInput) (*apimodel.PostsWithPagination, error) {
	s := dto.Session(ctx)
	if !s.IsAuthorized() {
		return nil, domain.NewError(domain.ErrCodeNotAuthorized)
	}

	posts, count, err := r.Posts.ByUserSubscriptions(ctx, dto.OptionsFromRest(o), s.User.ID)
	if err != nil {
		return nil, err
	}

	data := dto.PostsWithPaginationToRest(posts, *o, count)

	return &data, nil
}

func (r *queryResolver) ShowPost(ctx context.Context, id int64) (*apimodel.Post, error) {
	post, err := r.Posts.PostByID(ctx, id)
	if err != nil {
		return nil, err
	}

	data := dto.PostToRest(post)

	return &data, nil
}
