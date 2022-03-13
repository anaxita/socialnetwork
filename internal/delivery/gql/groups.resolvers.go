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

func (r *mutationResolver) AddGroup(ctx context.Context, g apimodel.AddGroupInput) (*entity.Group, error) {
	s := dto.Session(ctx)

	if !s.IsAuthorized() {
		return nil, domain.NewError(domain.ErrCodeNotAuthorized)
	}

	userPerms, err := r.Groups.UserPermissions(ctx, s.User.ID)
	if err != nil {
		return nil, err
	}

	ok := helpers.IsAccessAllowed(userPerms, domain.PermWrite)
	if !ok {
		return nil, domain.NewError(domain.ErrCodeForbidden)
	}

	group, err := r.Groups.Add(ctx, g.ToEntity())
	if err != nil {
		return nil, err
	}

	return &group, nil
}

func (r *mutationResolver) EditGroup(ctx context.Context, g apimodel.EditGroupInput) (*entity.Group, error) {
	s := dto.Session(ctx)

	if !s.IsAuthorized() {
		return nil, domain.NewError(domain.ErrCodeNotAuthorized)
	}

	userPerms, err := r.Groups.UserPermissions(ctx, s.User.ID, g.ID)
	if err != nil {
		return nil, err
	}

	ok := helpers.IsAccessAllowed(userPerms, domain.PermEdit)
	if !ok {
		return nil, domain.NewError(domain.ErrCodeForbidden)
	}

	editedGroup, err := r.Groups.Edit(ctx, g.ToEntity())
	if err != nil {
		return nil, err
	}

	return &editedGroup, nil
}

func (r *mutationResolver) DeleteGroup(ctx context.Context, id int64) (*bool, error) {
	s := dto.Session(ctx)

	if !s.IsAuthorized() {
		return nil, domain.NewError(domain.ErrCodeNotAuthorized)
	}

	userPerms, err := r.Groups.UserPermissions(ctx, s.User.ID, id)
	if err != nil {
		return nil, err
	}

	ok := helpers.IsAccessAllowed(userPerms, domain.PermDelete)
	if !ok {
		return nil, domain.NewError(domain.ErrCodeForbidden)
	}

	err = r.Groups.Remove(ctx, id)
	if err != nil {
		return nil, err
	}

	data := true

	return &data, nil
}

func (r *queryResolver) ShowGroups(ctx context.Context, o *apimodel.OptionsInput) (*apimodel.GroupsWithPagination, error) {
	o = helpers.ValidateOptions(o, apimodel.ColName)

	groups, count, err := r.Groups.ShowGroups(ctx, dto.OptionsFromRest(o))
	if err != nil {
		return nil, err
	}

	data := dto.GroupsWithPaginationToRest(groups, *o, count)

	return &data, nil
}

func (r *queryResolver) ShowGroup(ctx context.Context, groupID int64) (*entity.Group, error) {
	g, err := r.Groups.ShowGroup(ctx, groupID)
	if err != nil {
		return nil, err
	}

	return &g, nil
}
