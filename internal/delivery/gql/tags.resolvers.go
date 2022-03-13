package gql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"synergycommunity/internal/delivery/api/apimodel"
	"synergycommunity/internal/delivery/api/helpers"
	"synergycommunity/internal/domain"
	"synergycommunity/internal/domain/entity"
	"synergycommunity/internal/dto"
)

func (r *mutationResolver) AddTag(ctx context.Context, t apimodel.AddTagInput) (*entity.Tag, error) {
	s := dto.Session(ctx)

	if !s.IsAuthorized() {
		return nil, domain.NewError(domain.ErrCodeNotAuthorized)
	}

	userPerms, err := r.Groups.UserPermissions(ctx, s.User.ID)
	if err != nil {
		return nil, err
	}

	ok := helpers.IsAccessAllowed(userPerms, domain.PermEdit, domain.PermDelete,
		domain.PermAdministrate, domain.PermWrite)
	if !ok {
		return nil, domain.NewError(domain.ErrCodeForbidden)
	}

	tag, err := r.Tags.Create(ctx, dto.AddTagToEntity(t))
	if err != nil {
		return nil, err
	}

	return &tag, nil
}

func (r *mutationResolver) EditTag(ctx context.Context, t apimodel.EditTagInput) (*entity.Tag, error) {
	s := dto.Session(ctx)

	if !s.IsAuthorized() {
		return nil, domain.NewError(domain.ErrCodeNotAuthorized)
	}

	userPerms, err := r.Groups.UserPermissions(ctx, s.User.ID)
	if err != nil {
		return nil, err
	}

	ok := helpers.IsAccessAllowed(userPerms, domain.PermEdit, domain.PermDelete,
		domain.PermAdministrate, domain.PermWrite)
	if !ok {
		return nil, domain.NewError(domain.ErrCodeForbidden)
	}

	tag, err := r.Tags.Edit(ctx, dto.EditTagToEntity(t))
	if err != nil {
		return nil, err
	}

	return &tag, nil
}

func (r *mutationResolver) DeleteTag(ctx context.Context, id int64) (*bool, error) {
	s := dto.Session(ctx)

	if !s.IsAuthorized() {
		return nil, domain.NewError(domain.ErrCodeNotAuthorized)
	}

	userPerms, err := r.Groups.UserPermissions(ctx, s.User.ID)
	if err != nil {
		return nil, err
	}

	ok := helpers.IsAccessAllowed(userPerms, domain.PermEdit, domain.PermDelete,
		domain.PermAdministrate, domain.PermWrite)
	if !ok {
		return nil, domain.NewError(domain.ErrCodeForbidden)
	}

	err = r.Tags.Delete(ctx, id)
	if err != nil {
		return nil, err
	}

	data := true

	return &data, nil
}

func (r *queryResolver) ShowTags(ctx context.Context, o *apimodel.OptionsInput) (*apimodel.TagsWithPagination, error) {
	o = helpers.ValidateOptions(o, apimodel.ColID, apimodel.ColName)

	tag, count, err := r.Tags.Tags(ctx, dto.OptionsFromRest(o))
	if err != nil {
		return nil, err
	}

	data := dto.TagsWithPaginationToRest(tag, *o, count)

	return &data, nil
}

func (r *queryResolver) ShowTag(ctx context.Context, id int64) (*entity.Tag, error) {
	tag, err := r.Tags.TagByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return &tag, nil
}
