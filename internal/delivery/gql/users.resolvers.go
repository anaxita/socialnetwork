package gql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"synergycommunity/internal/delivery/api/apimodel"
	"synergycommunity/internal/delivery/api/helpers"
	"synergycommunity/internal/domain"
	"synergycommunity/internal/domain/entity"
	"synergycommunity/internal/dto"
)

func (r *mutationResolver) AddUser(ctx context.Context, u apimodel.UserInput) (*entity.User, error) {
	err := r.v.Struct(u)
	if err != nil {
		return nil, domain.NewValidationError(err)
	}

	user, err := r.Users.Register(ctx, dto.UserFromRest(u))
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *mutationResolver) EditProfile(ctx context.Context, u apimodel.UserInput) (*entity.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) EditUser(ctx context.Context, id int64, u apimodel.UserInput) (*entity.User, error) {
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

	err = r.v.Struct(u)
	if err != nil {
		return nil, domain.NewValidationError(err)
	}

	u.ID = id

	user, err := r.Users.Edit(ctx, dto.UserFromRest(u))
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *mutationResolver) DeleteUser(ctx context.Context, id int64) (*bool, error) {
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

	err = r.Users.Delete(ctx, id)
	if err != nil {
		return nil, err
	}

	ok = true

	return &ok, nil
}

func (r *mutationResolver) SetUserRole(ctx context.Context, ur apimodel.UserRole) (*bool, error) {
	s := dto.Session(ctx)
	if !s.IsAuthorized() {
		return nil, domain.NewError(domain.ErrCodeNotAuthorized)
	}

	err := r.v.Struct(ur)
	if err != nil {
		return nil, domain.NewErrorWrap(err, domain.ErrCodeBadRequest)
	}

	userPerms, err := r.Groups.UserPermissions(ctx, s.User.ID, ur.GroupID)
	if err != nil {
		return nil, err
	}

	isAllow := helpers.IsAccessAllowed(userPerms, domain.PermEdit, domain.PermDelete,
		domain.PermAdministrate, domain.PermWrite)
	if isAllow {
		return nil, domain.NewError(domain.ErrCodeForbidden)
	}

	err = r.Groups.SetUserRole(ctx, dto.UserRoleFromRest(ur))
	if err != nil {
		return nil, err
	}

	data := true

	return &data, nil
}

func (r *mutationResolver) UnsetUserRole(ctx context.Context, ur apimodel.UserRole) (*bool, error) {
	s := dto.Session(ctx)
	if !s.IsAuthorized() {
		return nil, domain.NewError(domain.ErrCodeNotAuthorized)
	}

	err := r.v.Struct(ur)
	if err != nil {
		return nil, domain.NewErrorWrap(err, domain.ErrCodeBadRequest)
	}

	userPerms, err := r.Groups.UserPermissions(ctx, s.User.ID, ur.GroupID)
	if err != nil {
		return nil, err
	}

	isAllow := helpers.IsAccessAllowed(userPerms, domain.PermEdit, domain.PermDelete,
		domain.PermAdministrate, domain.PermWrite)
	if isAllow {
		return nil, domain.NewError(domain.ErrCodeForbidden)
	}

	err = r.Groups.UnsetUserRole(ctx, dto.UserRoleFromRest(ur))
	if err != nil {
		return nil, err
	}

	data := true

	return &data, nil
}

func (r *mutationResolver) Subscribe(ctx context.Context, sub apimodel.UserSubscriptionInput) (*bool, error) {
	s := dto.Session(ctx)
	if !s.IsAuthorized() {
		return nil, domain.NewError(domain.ErrCodeNotAuthorized)
	}

	err := r.Users.Subscribe(ctx, sub.ToEntity(s.User.ID))
	if err != nil {
		return nil, err
	}

	data := true

	return &data, nil
}

func (r *mutationResolver) Unsubscribe(ctx context.Context, sub apimodel.UserSubscriptionInput) (*bool, error) {
	s := dto.Session(ctx)
	if !s.IsAuthorized() {
		return nil, domain.NewError(domain.ErrCodeNotAuthorized)
	}

	s.User.ID = 1

	err := r.Users.Unsubscribe(ctx, sub.ToEntity(s.User.ID))
	if err != nil {
		return nil, err
	}

	data := true

	return &data, nil
}

func (r *queryResolver) ShowUsers(ctx context.Context, o *apimodel.OptionsInput) (*apimodel.UsersWithPagination, error) {
	s := dto.Session(ctx)
	if !s.IsAuthorized() {
		return nil, domain.NewError(domain.ErrCodeNotAuthorized)
	}

	userPerms, err := r.Groups.UserPermissions(ctx, s.User.ID)
	if err != nil {
		return nil, err
	}

	isDisallow := helpers.IsDisallowAccess(userPerms, domain.PermNoWriting, domain.PermNoReading)
	if isDisallow {
		return nil, domain.NewError(domain.ErrCodeForbidden)
	}

	o = helpers.ValidateOptions(
		o, apimodel.ColID, apimodel.ColNickname, apimodel.ColFirstName,
		apimodel.ColLastName,
	)

	users, n, err := r.Users.Users(ctx, dto.OptionsFromRest(o))
	if err != nil {
		return nil, err
	}

	data := dto.UsersWithPaginationToRest(users, *o, n)

	return &data, nil
}

func (r *queryResolver) ShowUser(ctx context.Context, id int64) (*entity.User, error) {
	s := dto.Session(ctx)
	if !s.IsAuthorized() {
		return nil, domain.NewError(domain.ErrCodeNotAuthorized)
	}

	userPerms, err := r.Groups.UserPermissions(ctx, s.User.ID)
	if err != nil {
		return nil, err
	}

	isDisallow := helpers.IsDisallowAccess(userPerms, domain.PermNoWriting, domain.PermNoReading)
	if isDisallow {
		return nil, domain.NewError(domain.ErrCodeForbidden)
	}

	u, err := r.Users.UserByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return &u, nil
}

func (r *queryResolver) ShowSubscriptions(ctx context.Context) (*entity.UserSubscriptions, error) {
	s := dto.Session(ctx)
	if !s.IsAuthorized() {
		return nil, domain.NewError(domain.ErrCodeNotAuthorized)
	}

	subs, err := r.Users.Subscriptions(ctx, s.User.ID)
	if err != nil {
		return nil, err
	}

	return &subs, nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
